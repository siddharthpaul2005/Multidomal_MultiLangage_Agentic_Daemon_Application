import grpc
import time

from generated.manager_pb2 import RegisterAgentRequest
from generated.manager_pb2_grpc import ManagerStub

def main():
    channel = grpc.insecure_channel("localhost:50051")
    stub = ManagerStub(channel)

    response = stub.RegisterAgent(
        RegisterAgentRequest(
            agent_name="file_watcher_sigma_sus_boi",
            capabilities=["filesystem", "monitoring"]
        )
    )

    print("[Agent] Registered:", response.agent_id)

    while True:
        time.sleep(5)

if __name__ == "__main__":
    main()
