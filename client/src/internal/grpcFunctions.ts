import { grpc } from "@improbable-eng/grpc-web";
import { SapHistoryService } from "../internal/saphistory_pb_service";
import { Nothing, SystemList, SystemInfo, SapSelection, TransactionData } from "../internal/saphistory_pb";
import { SysInfo } from "./system";

// get systems from backend
export function GetSystems(grpcServer: string): Promise<any> {
  const nothing = new Nothing();

  const resPromise = new Promise<any>((resolve, reject) => {
    grpc.invoke(SapHistoryService.GetSystemList, {
      request: nothing,
      host: grpcServer,
      onMessage: (message: SystemList) => {
        resolve(message.toObject().systemsList);
      },
      onEnd: (
        code: grpc.Code,
        msg: string | undefined,
        trailers: grpc.Metadata
      ) => {
        if (code != grpc.Code.OK) {
          console.log("Error GetSystems: ", code, msg, trailers);
        }
        // if (code == grpc.Code.OK) {
        //   console.log("all ok");
        // } else {
        //   console.log("Error GetSystems: ", code, msg, trailers);
        // }
      },
    });
  });
  return resPromise;
}

// update database system entry
export function UpdateSystem(grpcServer: string, sys: SysInfo): Promise<any> {

  // !!!! besser waere, wenn keine zuordnung noetig ist und gleich der Typ SystemInfo uebergeben wird
  const system = new SystemInfo();
  system.setSid(sys.sid);
  system.setClient(sys.client);
  system.setDescription(sys.description);
  system.setSysnr(sys.sysnr);
  system.setHostname(sys.hostname);
  system.setUsername(sys.username);
  system.setPassword(sys.password);

  const resPromise = new Promise<any>((resolve, reject) => {
    grpc.invoke(SapHistoryService.UpdateSystem, {
      request: system,
      host: grpcServer,
      onEnd: (
        code: grpc.Code,
        msg: string | undefined,
        trailers: grpc.Metadata
      ) => {
        if (code != grpc.Code.OK) {
          console.log("Error UpdateSystem: ", code, msg, trailers);
        }
      },
    });
  });
  return resPromise;
}

// delete database system entry
export function DeleteSystem(grpcServer: string, sys: SysInfo): Promise<any> {

  // !!!! besser waere, wenn keine zuordnung noetig ist und gleich der Typ SystemInfo uebergeben wird
  const system = new SystemInfo();
  system.setSid(sys.sid);
  system.setClient(sys.client);
  system.setDescription(sys.description);
  system.setSysnr(sys.sysnr);
  system.setHostname(sys.hostname);
  system.setUsername(sys.username);
  system.setPassword(sys.password);

  const resPromise = new Promise<any>((resolve, reject) => {
    grpc.invoke(SapHistoryService.DeleteSystem, {
      request: system,
      host: grpcServer,
      onEnd: (
        code: grpc.Code,
        msg: string | undefined,
        trailers: grpc.Metadata
      ) => {
        if (code != grpc.Code.OK) {
          console.log("Error DeleteSystem: ", code, msg, trailers);
        }
      },
    });
  });
  return resPromise;
}


// get transaction data for system and timestamp
export function GetSapStatus(grpcHost: string, sid: string, ts: string, direction: number): Promise<any> {
  const sapSelection = new SapSelection();
  sapSelection.setSid(sid);
  sapSelection.setTs(ts);
  sapSelection.setDirection(direction);

  const resPromise = new Promise<any>((resolve, reject) => {
    grpc.invoke(SapHistoryService.GetSapStatus, {
      request: sapSelection,
      host: grpcHost,
      onMessage: (message: TransactionData) => {
        const data: any = message.toObject();
        let res: any = {};

        // !!!!! check ob head anzahl gleich data anzahl

        // headlines
        let headlines = new Array();
        let headline = new Array();

        // headlines
        for (let row of data.hdata.data2List) {
          row.data1List.unshift('#');
          row.data1List.forEach((col: any, i: number) => {
            let cols: any = {};
            cols["value"] = 'v' + String(i);
            cols["text"] = col;
            headline.push(cols);
          });
          headlines.push(headline);
          headline = [];
        }

        // data
        let tda = new Array();
        let td = new Array();
        for (let tData of data.tdata.data3List) {
          let line: number = 1;
          for (let row of tData.data2List) {
            let obj: any = {};
            obj['v0'] = line;
            line += 1;
            row.data1List.forEach((col: any, i: number) => {
              obj['v' + String(i + 1)] = col;
            });
            td.push(obj);
          }
          tda.push(td);
          td = []
        }


        res.ts = data.ts;
        res.tcvalues = data.tcodes.data1List

        res.head = headlines;
        res.data = tda;

        resolve(res);
      },
      onEnd: (
        code: grpc.Code,
        msg: string | undefined,
        trailers: grpc.Metadata
      ) => {
        if (code != grpc.Code.OK) {
          console.log("Error GetSapStatus: ", code, msg, trailers);
          // !!!!!!!!!!???????? reject()
        }
      },
    });
  });
  return resPromise;
}

function getArray(count: number): string[] {
  let arr: string[] = [];
  for (let i = 0; i < count; i++) {
    arr.push('v' + String(i));
  }
  return arr;
}