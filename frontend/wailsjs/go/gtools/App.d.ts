// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {internal} from '../models';
import {util} from '../models';

export function AddMdDirPath(arg1:internal.MdDir):Promise<util.Resp>;

export function AddTodoItem(arg1:internal.TodoItem):Promise<util.Resp>;

export function DelMdDir(arg1:internal.MdDir):Promise<util.Resp>;

export function DelTodoItem(arg1:internal.TodoItem):Promise<util.Resp>;

export function GetMdContent(arg1:string):Promise<util.Resp>;

export function GetMdDirList():Promise<util.Resp>;

export function GetMdFileList(arg1:string):Promise<util.Resp>;

export function GetTodoItem(arg1:number):Promise<util.Resp>;

export function GetTodoList():Promise<util.Resp>;

export function NewMd(arg1:string,arg2:string):Promise<util.Resp>;

export function OpenMdFolderWindow():Promise<util.Resp>;

export function OpenMdSaveFileWindow():Promise<util.Resp>;

export function OpenSelectFileWindow():Promise<util.Resp>;

export function SaveMdContent(arg1:string,arg2:string):Promise<util.Resp>;

export function StartTomcat():Promise<string>;

export function UpdateTodoItem(arg1:internal.TodoItem):Promise<util.Resp>;

export function UploadScreenshot():Promise<util.Resp>;
