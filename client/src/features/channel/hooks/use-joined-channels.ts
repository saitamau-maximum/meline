import { useQuery } from "@tanstack/react-query";
import { ChannelRepositoryImpl } from "@/repositories/channel";
import { useMemo } from "react";

export const useJoinedChannels = () => {
  const channelRepository = useMemo(() => new ChannelRepositoryImpl(), []);

  return useQuery({
    queryKey: channelRepository.getJoinedChannels$$key(),
    queryFn: () => channelRepository.getJoinedChannels(),
  });
};
