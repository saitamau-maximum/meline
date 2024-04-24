import { useQuery } from "@tanstack/react-query";
import { ChannelRepositoryImpl } from "@/repositories/channel";

export const useJoinedChannels = () => {
  return useQuery({
    queryKey: ChannelRepositoryImpl.getJoinedChannels$$key(),
    queryFn: () => ChannelRepositoryImpl.getJoinedChannels(),
  });
};
