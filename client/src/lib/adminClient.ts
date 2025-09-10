import { GrpcWebFetchTransport } from "@protobuf-ts/grpcweb-transport";
import { AdminServiceClient } from "../proto/Admin.client";

const transport = new GrpcWebFetchTransport({
    baseUrl: import.meta.env.VITE_ENV === "development" ? import.meta.env.VITE_ENVOY_URL : '/rpc/s/'
})
const client = new AdminServiceClient(transport);

export default client