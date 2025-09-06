import { GrpcWebFetchTransport } from "@protobuf-ts/grpcweb-transport";
import { SmsServicesClient } from "../proto/SmsApi.client";

const transport = new GrpcWebFetchTransport({
    baseUrl: import.meta.env.VITE_ENVOY_URL,
})
const client = new SmsServicesClient(transport);

export default client