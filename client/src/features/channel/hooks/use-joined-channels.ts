import { useQuery } from "@tanstack/react-query";
import { useRepositories } from "@/hooks/repository";

export const useJoinedChannels = () => {
  const { channelRepository } = useRepositories();

  return useQuery({
    queryKey: channelRepository.getJoinedChannels$$key(),
    queryFn: () => channelRepository.getJoinedChannels(),
  });
};
