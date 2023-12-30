import { IChannelRepository } from "@/repositories/channel";
import { valibotResolver } from "@hookform/resolvers/valibot";
import { useState } from "react";
import { useForm } from "react-hook-form";
import { Input, maxLength, minLength, object, string } from "valibot";

const CreateChannelSchema = object({
  name: string([
    minLength(1, "チャンネル名は必須です。"),
    maxLength(255, "チャンネル名は255文字以内で入力してください。"),
  ]),
});

type CreateChannelFormData = Input<typeof CreateChannelSchema>;

interface CreateChannelFormProps {
  channelRepository: IChannelRepository;
  onCreated: () => void;
}

export const useCreateChannelForm = ({
  channelRepository,
  onCreated,
}: CreateChannelFormProps) => {
  const {
    reset,
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<CreateChannelFormData>({
    resolver: valibotResolver(CreateChannelSchema),
  });
  const [formError, setFormError] = useState("");

  const onSubmit = async (form: CreateChannelFormData) => {
    const res = await channelRepository.createChannel(form.name);

    if (res.ok) {
      setFormError("");
      onCreated();
      return;
    }

    if (res.status >= 400 && res.status < 500) {
      const { message } = (await res.json()) as { message: string };
      setFormError(message);
      return;
    }

    setFormError("エラーが発生しました、時間を空けてもう一度お試しください。");
  };

  return {
    reset: () => {
      reset();
      setFormError("");
    },
    register,
    handleSubmit: handleSubmit(onSubmit),
    errors,
    formError,
  };
};
