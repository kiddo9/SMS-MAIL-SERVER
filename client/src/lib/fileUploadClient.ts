import { GrpcWebFetchTransport } from "@protobuf-ts/grpcweb-transport";
import { fileUploadServicesClient } from "../proto/FileUpload.client";


const transport = new GrpcWebFetchTransport({
    baseUrl: '/rpc/s/'
})

const client = new fileUploadServicesClient(transport);

export default client