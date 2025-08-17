import * as jspb from 'google-protobuf'



export class Admin extends jspb.Message {
  getName(): string;
  setName(value: string): Admin;

  getEmail(): string;
  setEmail(value: string): Admin;

  getRole(): string;
  setRole(value: string): Admin;

  getPhone(): string;
  setPhone(value: string): Admin;

  getOtp(): number;
  setOtp(value: number): Admin;

  getOtpexpiry(): number;
  setOtpexpiry(value: number): Admin;

  getApikey(): string;
  setApikey(value: string): Admin;

  getUuid(): string;
  setUuid(value: string): Admin;

  getJwt(): string;
  setJwt(value: string): Admin;

  getEmailverified(): boolean;
  setEmailverified(value: boolean): Admin;

  getCreatedat(): string;
  setCreatedat(value: string): Admin;

  getUpdatedat(): string;
  setUpdatedat(value: string): Admin;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Admin.AsObject;
  static toObject(includeInstance: boolean, msg: Admin): Admin.AsObject;
  static serializeBinaryToWriter(message: Admin, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Admin;
  static deserializeBinaryFromReader(message: Admin, reader: jspb.BinaryReader): Admin;
}

export namespace Admin {
  export type AsObject = {
    name: string,
    email: string,
    role: string,
    phone: string,
    otp: number,
    otpexpiry: number,
    apikey: string,
    uuid: string,
    jwt: string,
    emailverified: boolean,
    createdat: string,
    updatedat: string,
  }
}

export class CreateAdminUserRequest extends jspb.Message {
  getRecaptchatoken(): string;
  setRecaptchatoken(value: string): CreateAdminUserRequest;

  getAdmin(): Admin | undefined;
  setAdmin(value?: Admin): CreateAdminUserRequest;
  hasAdmin(): boolean;
  clearAdmin(): CreateAdminUserRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateAdminUserRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateAdminUserRequest): CreateAdminUserRequest.AsObject;
  static serializeBinaryToWriter(message: CreateAdminUserRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateAdminUserRequest;
  static deserializeBinaryFromReader(message: CreateAdminUserRequest, reader: jspb.BinaryReader): CreateAdminUserRequest;
}

export namespace CreateAdminUserRequest {
  export type AsObject = {
    recaptchatoken: string,
    admin?: Admin.AsObject,
  }
}

export class CreateAdminUserResponse extends jspb.Message {
  getAdmincreated(): boolean;
  setAdmincreated(value: boolean): CreateAdminUserResponse;

  getMessage(): string;
  setMessage(value: string): CreateAdminUserResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateAdminUserResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateAdminUserResponse): CreateAdminUserResponse.AsObject;
  static serializeBinaryToWriter(message: CreateAdminUserResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateAdminUserResponse;
  static deserializeBinaryFromReader(message: CreateAdminUserResponse, reader: jspb.BinaryReader): CreateAdminUserResponse;
}

export namespace CreateAdminUserResponse {
  export type AsObject = {
    admincreated: boolean,
    message: string,
  }
}

export class GetAndValidateAdminRequest extends jspb.Message {
  getEmail(): string;
  setEmail(value: string): GetAndValidateAdminRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetAndValidateAdminRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetAndValidateAdminRequest): GetAndValidateAdminRequest.AsObject;
  static serializeBinaryToWriter(message: GetAndValidateAdminRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetAndValidateAdminRequest;
  static deserializeBinaryFromReader(message: GetAndValidateAdminRequest, reader: jspb.BinaryReader): GetAndValidateAdminRequest;
}

export namespace GetAndValidateAdminRequest {
  export type AsObject = {
    email: string,
  }
}

export class ValidateAdminResponse extends jspb.Message {
  getIsvalid(): boolean;
  setIsvalid(value: boolean): ValidateAdminResponse;

  getAdmin(): Admin | undefined;
  setAdmin(value?: Admin): ValidateAdminResponse;
  hasAdmin(): boolean;
  clearAdmin(): ValidateAdminResponse;

  getIsemailverified(): boolean;
  setIsemailverified(value: boolean): ValidateAdminResponse;

  getMessage(): string;
  setMessage(value: string): ValidateAdminResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ValidateAdminResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ValidateAdminResponse): ValidateAdminResponse.AsObject;
  static serializeBinaryToWriter(message: ValidateAdminResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ValidateAdminResponse;
  static deserializeBinaryFromReader(message: ValidateAdminResponse, reader: jspb.BinaryReader): ValidateAdminResponse;
}

export namespace ValidateAdminResponse {
  export type AsObject = {
    isvalid: boolean,
    admin?: Admin.AsObject,
    isemailverified: boolean,
    message: string,
  }
}

export class AllAdmins extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AllAdmins.AsObject;
  static toObject(includeInstance: boolean, msg: AllAdmins): AllAdmins.AsObject;
  static serializeBinaryToWriter(message: AllAdmins, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AllAdmins;
  static deserializeBinaryFromReader(message: AllAdmins, reader: jspb.BinaryReader): AllAdmins;
}

export namespace AllAdmins {
  export type AsObject = {
  }
}

export class OtpRequest extends jspb.Message {
  getEmail(): string;
  setEmail(value: string): OtpRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): OtpRequest.AsObject;
  static toObject(includeInstance: boolean, msg: OtpRequest): OtpRequest.AsObject;
  static serializeBinaryToWriter(message: OtpRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): OtpRequest;
  static deserializeBinaryFromReader(message: OtpRequest, reader: jspb.BinaryReader): OtpRequest;
}

export namespace OtpRequest {
  export type AsObject = {
    email: string,
  }
}

export class OtpResponse extends jspb.Message {
  getOtpsent(): boolean;
  setOtpsent(value: boolean): OtpResponse;

  getMessage(): string;
  setMessage(value: string): OtpResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): OtpResponse.AsObject;
  static toObject(includeInstance: boolean, msg: OtpResponse): OtpResponse.AsObject;
  static serializeBinaryToWriter(message: OtpResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): OtpResponse;
  static deserializeBinaryFromReader(message: OtpResponse, reader: jspb.BinaryReader): OtpResponse;
}

export namespace OtpResponse {
  export type AsObject = {
    otpsent: boolean,
    message: string,
  }
}

export class OtpVerificationRequest extends jspb.Message {
  getEmail(): string;
  setEmail(value: string): OtpVerificationRequest;

  getOtp(): string;
  setOtp(value: string): OtpVerificationRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): OtpVerificationRequest.AsObject;
  static toObject(includeInstance: boolean, msg: OtpVerificationRequest): OtpVerificationRequest.AsObject;
  static serializeBinaryToWriter(message: OtpVerificationRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): OtpVerificationRequest;
  static deserializeBinaryFromReader(message: OtpVerificationRequest, reader: jspb.BinaryReader): OtpVerificationRequest;
}

export namespace OtpVerificationRequest {
  export type AsObject = {
    email: string,
    otp: string,
  }
}

export class OtpVerificationResponse extends jspb.Message {
  getIsverified(): boolean;
  setIsverified(value: boolean): OtpVerificationResponse;

  getMessage(): string;
  setMessage(value: string): OtpVerificationResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): OtpVerificationResponse.AsObject;
  static toObject(includeInstance: boolean, msg: OtpVerificationResponse): OtpVerificationResponse.AsObject;
  static serializeBinaryToWriter(message: OtpVerificationResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): OtpVerificationResponse;
  static deserializeBinaryFromReader(message: OtpVerificationResponse, reader: jspb.BinaryReader): OtpVerificationResponse;
}

export namespace OtpVerificationResponse {
  export type AsObject = {
    isverified: boolean,
    message: string,
  }
}

export class GetASpecificAdminRequest extends jspb.Message {
  getUuid(): string;
  setUuid(value: string): GetASpecificAdminRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetASpecificAdminRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetASpecificAdminRequest): GetASpecificAdminRequest.AsObject;
  static serializeBinaryToWriter(message: GetASpecificAdminRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetASpecificAdminRequest;
  static deserializeBinaryFromReader(message: GetASpecificAdminRequest, reader: jspb.BinaryReader): GetASpecificAdminRequest;
}

export namespace GetASpecificAdminRequest {
  export type AsObject = {
    uuid: string,
  }
}

export class GetASpecificAdminResponse extends jspb.Message {
  getAdmin(): Admin | undefined;
  setAdmin(value?: Admin): GetASpecificAdminResponse;
  hasAdmin(): boolean;
  clearAdmin(): GetASpecificAdminResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetASpecificAdminResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetASpecificAdminResponse): GetASpecificAdminResponse.AsObject;
  static serializeBinaryToWriter(message: GetASpecificAdminResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetASpecificAdminResponse;
  static deserializeBinaryFromReader(message: GetASpecificAdminResponse, reader: jspb.BinaryReader): GetASpecificAdminResponse;
}

export namespace GetASpecificAdminResponse {
  export type AsObject = {
    admin?: Admin.AsObject,
  }
}

export class AdminUpdateRequest extends jspb.Message {
  getAdmin(): Admin | undefined;
  setAdmin(value?: Admin): AdminUpdateRequest;
  hasAdmin(): boolean;
  clearAdmin(): AdminUpdateRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AdminUpdateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AdminUpdateRequest): AdminUpdateRequest.AsObject;
  static serializeBinaryToWriter(message: AdminUpdateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AdminUpdateRequest;
  static deserializeBinaryFromReader(message: AdminUpdateRequest, reader: jspb.BinaryReader): AdminUpdateRequest;
}

export namespace AdminUpdateRequest {
  export type AsObject = {
    admin?: Admin.AsObject,
  }
}

export class AdminUpdateResponse extends jspb.Message {
  getIsupdated(): boolean;
  setIsupdated(value: boolean): AdminUpdateResponse;

  getMessage(): string;
  setMessage(value: string): AdminUpdateResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AdminUpdateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: AdminUpdateResponse): AdminUpdateResponse.AsObject;
  static serializeBinaryToWriter(message: AdminUpdateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AdminUpdateResponse;
  static deserializeBinaryFromReader(message: AdminUpdateResponse, reader: jspb.BinaryReader): AdminUpdateResponse;
}

export namespace AdminUpdateResponse {
  export type AsObject = {
    isupdated: boolean,
    message: string,
  }
}

export class DeleteAdminRequest extends jspb.Message {
  getUuid(): string;
  setUuid(value: string): DeleteAdminRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteAdminRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteAdminRequest): DeleteAdminRequest.AsObject;
  static serializeBinaryToWriter(message: DeleteAdminRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteAdminRequest;
  static deserializeBinaryFromReader(message: DeleteAdminRequest, reader: jspb.BinaryReader): DeleteAdminRequest;
}

export namespace DeleteAdminRequest {
  export type AsObject = {
    uuid: string,
  }
}

export class DeleteAdminResponse extends jspb.Message {
  getIsdeleted(): boolean;
  setIsdeleted(value: boolean): DeleteAdminResponse;

  getMessage(): string;
  setMessage(value: string): DeleteAdminResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteAdminResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteAdminResponse): DeleteAdminResponse.AsObject;
  static serializeBinaryToWriter(message: DeleteAdminResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteAdminResponse;
  static deserializeBinaryFromReader(message: DeleteAdminResponse, reader: jspb.BinaryReader): DeleteAdminResponse;
}

export namespace DeleteAdminResponse {
  export type AsObject = {
    isdeleted: boolean,
    message: string,
  }
}

export class TokenValidationRequest extends jspb.Message {
  getToken(): string;
  setToken(value: string): TokenValidationRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TokenValidationRequest.AsObject;
  static toObject(includeInstance: boolean, msg: TokenValidationRequest): TokenValidationRequest.AsObject;
  static serializeBinaryToWriter(message: TokenValidationRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TokenValidationRequest;
  static deserializeBinaryFromReader(message: TokenValidationRequest, reader: jspb.BinaryReader): TokenValidationRequest;
}

export namespace TokenValidationRequest {
  export type AsObject = {
    token: string,
  }
}

export class TokenValidationResponse extends jspb.Message {
  getIsvalid(): boolean;
  setIsvalid(value: boolean): TokenValidationResponse;

  getEmail(): string;
  setEmail(value: string): TokenValidationResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TokenValidationResponse.AsObject;
  static toObject(includeInstance: boolean, msg: TokenValidationResponse): TokenValidationResponse.AsObject;
  static serializeBinaryToWriter(message: TokenValidationResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TokenValidationResponse;
  static deserializeBinaryFromReader(message: TokenValidationResponse, reader: jspb.BinaryReader): TokenValidationResponse;
}

export namespace TokenValidationResponse {
  export type AsObject = {
    isvalid: boolean,
    email: string,
  }
}

