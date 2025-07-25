// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {command} from '../models';
import {commandgroup} from '../models';
import {config} from '../models';

export function AddCommand(arg1:command.Command):Promise<void>;

export function EditCommand(arg1:command.Command):Promise<void>;

export function GetCommandGroups():Promise<Array<commandgroup.CommandGroup>>;

export function GetCommands():Promise<Record<string, command.Command>>;

export function GetUserConfig():Promise<config.UserConfig>;

export function RemoveCommand(arg1:string):Promise<void>;

export function RunCommand(arg1:string):Promise<Record<string, command.Command>>;

export function SaveCommandGroups(arg1:Array<commandgroup.CommandGroup>):Promise<void>;

export function SaveUserConfig(arg1:config.UserConfig):Promise<void>;

export function StopCommand(arg1:string):Promise<void>;
