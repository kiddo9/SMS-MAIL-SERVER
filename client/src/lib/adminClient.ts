import { GrpcWebFetchTransport } from "@protobuf-ts/grpcweb-transport";
import { AdminServiceClient } from "../proto/Admin.client";

const transport = new GrpcWebFetchTransport({
    baseUrl: '/rpc/s/'
})
const client = new AdminServiceClient(transport);

export default client