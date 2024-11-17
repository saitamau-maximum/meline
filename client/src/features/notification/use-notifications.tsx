import { useRepositories } from "@/hooks/repository";
import { logger } from "@/utils/logger";
import { useEffect } from "react";
import { useLocation } from "react-router-dom";
import { Message } from "./components/message";
import { useToast } from "@/components/ui/toast/use-toast";

const channelRegex = /^\/channel\/(?<channelId>\d+)$/;

export const useNotifications = () => {
  const { notificationRepository } = useRepositories();
  const { pushToast } = useToast();
  const location = useLocation();

  useEffect(() => {
    logger.raw(`=== useSetup ===`);
    notificationRepository.connect();
    logger.info(`Connecting to notification channel`);
    notificationRepository.onMessageReceived((res) => {
      const channelMatch = location.pathname.match(channelRegex);
      if (
        channelMatch &&
        channelMatch.groups?.channelId === res.payload.channel_id.toString()
      ) {
        // 現在のチャンネルに関連する通知の場合は何もしない
        return;
      }
      pushToast(
        <Message
          avatarUrl={res.payload.message.user.image_url}
          username={res.payload.message.user.name}
          message={res.payload.message.content}
        />,
        {
          to: `/channel/${res.payload.channel_id}`,
        }
      );
    });

    return () => {
      notificationRepository.disconnect();
      logger.info(`Disconnecting from notification channel`);
    };
  }, [notificationRepository, pushToast, location]);
};
