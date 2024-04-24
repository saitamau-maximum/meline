import { useMutation, useQueryClient } from "@tanstack/react-query";
import { ChannelRepositoryImpl } from "@/repositories/channel";

interface UseCreateChannelsOptions {
  onCreated: () => void;
  onFailed: () => void;
}

interface CreateChannelParams {
  name: string;
}

export const useCreateChannels = ({
  onCreated,
  onFailed,
}: UseCreateChannelsOptions) => {
  const client = useQueryClient();

  return useMutation({
    mutationFn: async ({ name }: CreateChannelParams) => {
      return ChannelRepositoryImpl.createChannel(name);
    },
    onSettled: () => {
      client.invalidateQueries({
        queryKey: ChannelRepositoryImpl.getJoinedChannels$$key(),
      });
    },
    onSuccess: () => {
      onCreated();
    },
    onError: () => {
      onFailed();
    },
  });
};
