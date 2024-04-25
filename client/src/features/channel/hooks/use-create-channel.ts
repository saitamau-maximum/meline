import { useMutation, useQueryClient } from "@tanstack/react-query";
import { ChannelRepositoryImpl } from "@/repositories/channel";

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

  return useMutation({
    mutationFn: async (param: MutationParam) => {
      return ChannelRepositoryImpl.createChannel(param);
    },
    onSettled: () => {
      client.invalidateQueries({
        queryKey: ChannelRepositoryImpl.getJoinedChannels$$key(),
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
