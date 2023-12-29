import React, { createContext, useCallback, useState } from "react";
import { object, string, safeParse, number, array } from "valibot";

const ChannelListResponse = object({
  channels: array(
    object({
      id: number(),
      name: string(),
    })
  ),
});

type Channel = {
  id: number;
  name: string;
};

type ChannelContextProps = {
  joinedChannels: Channel[];
  fetchJoinedChannels: () => Promise<void>;
};

export const ChannelContext = createContext<ChannelContextProps>({
  joinedChannels: [],
  // eslint-disable-next-line @typescript-eslint/no-empty-function
  fetchJoinedChannels: async () => {},
});

interface ChannelProviderProps {
  children: React.ReactNode;
}

export const ChannelProvider = ({ children }: ChannelProviderProps) => {
  const [joinedChannels, setJoinedChannels] = useState<Channel[]>([]);

  const fetchJoinedChannels = useCallback(async () => {
    const res = await fetch("/api/channels");
    if (!res.ok) return setJoinedChannels([]);
    const validated = safeParse(ChannelListResponse, await res.json());
    if (!validated.success) return setJoinedChannels([]);
    setJoinedChannels(validated.output.channels);
  }, [setJoinedChannels]);

  return (
    <ChannelContext.Provider
      value={{
        joinedChannels,
        fetchJoinedChannels,
      }}
    >
      {children}
    </ChannelContext.Provider>
  );
};
