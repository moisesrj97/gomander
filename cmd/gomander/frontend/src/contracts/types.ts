import {
  command,
  commandgroup,
  config,
  event,
} from "../../wailsjs/go/models.ts";

// Types
export type Command = command.Command;
export type UserConfig = config.UserConfig;
export type CommandGroup = commandgroup.CommandGroup;

// Enums
export enum Event {
  GET_COMMANDS = event.Event.GET_COMMANDS,
  NEW_LOG_ENTRY = event.Event.NEW_LOG_ENTRY,
  SUCCESS_NOTIFICATION = event.Event.SUCCESS_NOTIFICATION,
  ERROR_NOTIFICATION = event.Event.ERROR_NOTIFICATION,
  PROCESS_FINISHED = event.Event.PROCESS_FINISHED,
  GET_USER_CONFIG = event.Event.GET_USER_CONFIG,
  GET_COMMAND_GROUPS = event.Event.GET_COMMAND_GROUPS,
}

export type EventData = {
  [Event.GET_COMMANDS]: null;
  [Event.NEW_LOG_ENTRY]: {
    id: string;
    line: string;
  };
  [Event.ERROR_NOTIFICATION]: string;
  [Event.PROCESS_FINISHED]: string;
  [Event.SUCCESS_NOTIFICATION]: string;
  [Event.GET_USER_CONFIG]: null;
  [Event.GET_COMMAND_GROUPS]: null;
};
