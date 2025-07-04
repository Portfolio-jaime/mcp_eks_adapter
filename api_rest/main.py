from fastapi import FastAPI
import grpc
import os
import sys

dir_actual = os.path.dirname(os.path.abspath(__file__))
sys.path.append(os.path.abspath(os.path.join(dir_actual, '../client')))
import context_pb2, context_pb2_grpc

app = FastAPI()
GRPC_SERVER = "localhost:50051"

@app.get("/kubernetes/version")
def get_kubernetes_version():
    with grpc.insecure_channel(GRPC_SERVER) as channel:
        stub = context_pb2_grpc.ModelContextServiceStub(channel)
        response = stub.GetContext(context_pb2.ContextRequest(cluster_name="default"))
        return {"kubernetes_version": response.kubernetes_version}

@app.get("/helm/charts")
def get_helm_charts():
    with grpc.insecure_channel(GRPC_SERVER) as channel:
        stub = context_pb2_grpc.ModelContextServiceStub(channel)
        response = stub.GetContext(context_pb2.ContextRequest(cluster_name="default"))
        charts = [
            {"name": c.name, "version": c.version, "namespace": c.namespace}
            for c in response.charts
        ]
        return {"charts": charts}
