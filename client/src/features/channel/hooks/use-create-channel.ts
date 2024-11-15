import { useMutation, useQueryClient } from "@tanstack/react-query";
import { ChannelRepositoryImpl } from "@/repositories/channel";
import { useMemo } from "react";

interface UseCreateChannelsOptions {
  onCreated?: () => void;
  onFailed?: () => void;
}

interface MutationParam {
  name: string;
}

export const useCreateChannels = ({
  onCreated,
  onFailed,
}: UseCreateChannelsOptions) => {
  const client = useQueryClient();
  const channelRepository = useMemo(() => new ChannelRepositoryImpl(), []);

  return useMutation({
    mutationFn: async (param: MutationParam) => {
      return channelRepository.createChannel(param);
    },
    onSettled: () => {
      client.invalidateQueries({
        queryKey: channelRepository.getJoinedChannels$$key(),
      });
    },
    onSuccess: () => {
      onCreated?.();
    },
    onError: () => {
      onFailed?.();
    },
  });
};
