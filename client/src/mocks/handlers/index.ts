import { userHandlers } from "./user";
import { channelHandlers } from "./channel";

export const handlers = [...userHandlers, ...channelHandlers];
