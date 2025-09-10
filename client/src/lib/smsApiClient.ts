import { GrpcWebFetchTransport } from "@protobuf-ts/grpcweb-transport";
import { SmsServicesClient } from "../proto/SmsApi.client";

const transport = new GrpcWebFetchTransport({
    baseUrl: import.meta.env.VITE_ENV === "development" ? import.meta.env.VITE_ENVOY_URL : '/rpc/s/'
})
const client = new SmsServicesClient(transport);

export default client