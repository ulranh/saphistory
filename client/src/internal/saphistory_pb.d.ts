// package: private
// file: internal/saphistory.proto

import * as jspb from "google-protobuf";

export class Secret extends jspb.Message {
  getNameMap(): jspb.Map<string, Uint8Array | string>;
  clearNameMap(): void;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Secret.AsObject;
  static toObject(includeInstance: boolean, msg: Secret): Secret.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Secret, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Secret;
  static deserializeBinaryFromReader(message: Secret, reader: jspb.BinaryReader): Secret;
}

export namespace Secret {
  export type AsObject = {
    nameMap: Array<[string, Uint8Array | string]>,
  }
}

export class D1StringList extends jspb.Message {
  clearData1List(): void;
  getData1List(): Array<string>;
  setData1List(value: Array<string>): void;
  addData1(value: string, index?: number): string;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): D1StringList.AsObject;
  static toObject(includeInstance: boolean, msg: D1StringList): D1StringList.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: D1StringList, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): D1StringList;
  static deserializeBinaryFromReader(message: D1StringList, reader: jspb.BinaryReader): D1StringList;
}

export namespace D1StringList {
  export type AsObject = {
    data1List: Array<string>,
  }
}

export class D2StringList extends jspb.Message {
  clearData2List(): void;
  getData2List(): Array<D1StringList>;
  setData2List(value: Array<D1StringList>): void;
  addData2(value?: D1StringList, index?: number): D1StringList;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): D2StringList.AsObject;
  static toObject(includeInstance: boolean, msg: D2StringList): D2StringList.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: D2StringList, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): D2StringList;
  static deserializeBinaryFromReader(message: D2StringList, reader: jspb.BinaryReader): D2StringList;
}

export namespace D2StringList {
  export type AsObject = {
    data2List: Array<D1StringList.AsObject>,
  }
}

export class D3StringList extends jspb.Message {
  clearData3List(): void;
  getData3List(): Array<D2StringList>;
  setData3List(value: Array<D2StringList>): void;
  addData3(value?: D2StringList, index?: number): D2StringList;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): D3StringList.AsObject;
  static toObject(includeInstance: boolean, msg: D3StringList): D3StringList.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: D3StringList, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): D3StringList;
  static deserializeBinaryFromReader(message: D3StringList, reader: jspb.BinaryReader): D3StringList;
}

export namespace D3StringList {
  export type AsObject = {
    data3List: Array<D2StringList.AsObject>,
  }
}

export class TransactionData extends jspb.Message {
  getTs(): string;
  setTs(value: string): void;

  hasTcodes(): boolean;
  clearTcodes(): void;
  getTcodes(): D1StringList | undefined;
  setTcodes(value?: D1StringList): void;

  hasHdata(): boolean;
  clearHdata(): void;
  getHdata(): D2StringList | undefined;
  setHdata(value?: D2StringList): void;

  hasTdata(): boolean;
  clearTdata(): void;
  getTdata(): D3StringList | undefined;
  setTdata(value?: D3StringList): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TransactionData.AsObject;
  static toObject(includeInstance: boolean, msg: TransactionData): TransactionData.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: TransactionData, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TransactionData;
  static deserializeBinaryFromReader(message: TransactionData, reader: jspb.BinaryReader): TransactionData;
}

export namespace TransactionData {
  export type AsObject = {
    ts: string,
    tcodes?: D1StringList.AsObject,
    hdata?: D2StringList.AsObject,
    tdata?: D3StringList.AsObject,
  }
}

export class SapSelection extends jspb.Message {
  getSid(): string;
  setSid(value: string): void;

  getTs(): string;
  setTs(value: string): void;

  getDirection(): number;
  setDirection(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SapSelection.AsObject;
  static toObject(includeInstance: boolean, msg: SapSelection): SapSelection.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SapSelection, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SapSelection;
  static deserializeBinaryFromReader(message: SapSelection, reader: jspb.BinaryReader): SapSelection;
}

export namespace SapSelection {
  export type AsObject = {
    sid: string,
    ts: string,
    direction: number,
  }
}

export class SystemInfo extends jspb.Message {
  getSid(): string;
  setSid(value: string): void;

  getClient(): string;
  setClient(value: string): void;

  getDescription(): string;
  setDescription(value: string): void;

  getSysnr(): string;
  setSysnr(value: string): void;

  getHostname(): string;
  setHostname(value: string): void;

  getUsername(): string;
  setUsername(value: string): void;

  getPassword(): string;
  setPassword(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SystemInfo.AsObject;
  static toObject(includeInstance: boolean, msg: SystemInfo): SystemInfo.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SystemInfo, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SystemInfo;
  static deserializeBinaryFromReader(message: SystemInfo, reader: jspb.BinaryReader): SystemInfo;
}

export namespace SystemInfo {
  export type AsObject = {
    sid: string,
    client: string,
    description: string,
    sysnr: string,
    hostname: string,
    username: string,
    password: string,
  }
}

export class SystemList extends jspb.Message {
  clearSystemsList(): void;
  getSystemsList(): Array<SystemInfo>;
  setSystemsList(value: Array<SystemInfo>): void;
  addSystems(value?: SystemInfo, index?: number): SystemInfo;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SystemList.AsObject;
  static toObject(includeInstance: boolean, msg: SystemList): SystemList.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SystemList, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SystemList;
  static deserializeBinaryFromReader(message: SystemList, reader: jspb.BinaryReader): SystemList;
}

export namespace SystemList {
  export type AsObject = {
    systemsList: Array<SystemInfo.AsObject>,
  }
}

export class Nothing extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Nothing.AsObject;
  static toObject(includeInstance: boolean, msg: Nothing): Nothing.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Nothing, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Nothing;
  static deserializeBinaryFromReader(message: Nothing, reader: jspb.BinaryReader): Nothing;
}

export namespace Nothing {
  export type AsObject = {
  }
}

