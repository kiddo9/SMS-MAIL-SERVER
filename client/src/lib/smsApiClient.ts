import { GrpcWebFetchTransport } from "@protobuf-ts/grpcweb-transport";
import { SmsServicesClient } from "../proto/SmsApi.client";

const transport = new GrpcWebFetchTransport({
    baseUrl: '/rpc/s/',
})
const client = new SmsServicesClient(transport);

export default client