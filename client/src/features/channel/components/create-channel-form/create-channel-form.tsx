import { Dialog } from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
import { useCallback, useState } from "react";
import { TextInput } from "@/components/ui/text-input";
import { styles } from "./create-channel-form.css";
import { useForm } from "react-hook-form";
import { valibotResolver } from "@hookform/resolvers/valibot";
import { Input, maxLength, minLength, object, string } from "valibot";
import { useCreateChannels } from "../../hooks/use-create-channel";

const CreateChannelSchema = object({
  name: string([
    minLength(1, "チャンネル名は必須です。"),
    maxLength(255, "チャンネル名は255文字以内で入力してください。"),
  ]),
});

type CreateChannelFormData = Input<typeof CreateChannelSchema>;

export const CreateChannelForm = () => {
  const {
    reset,
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<CreateChannelFormData>({
    resolver: valibotResolver(CreateChannelSchema),
  });
  const [isOpen, setIsOpen] = useState(false);

  const open = useCallback(() => {
    setIsOpen(true);
  }, []);

  const close = useCallback(() => {
    reset();
    setIsOpen(false);
  }, [reset]);

  const [formError, setFormError] = useState("");

  const handleCreated = useCallback(() => {
    close();
  }, [close]);

  const handleFailed = useCallback(() => {
    setFormError("チャンネルの作成に失敗しました。");
  }, []);

  const { mutate: createChannel } = useCreateChannels({
    onCreated: handleCreated,
    onFailed: handleFailed,
  });

  const onSubmit = useCallback(
    (data: CreateChannelFormData) => {
      createChannel(data);
    },
    [createChannel]
  );

  return (
    <Dialog.Root
      open={isOpen}
      onOpenChange={(isOpen) => {
        isOpen ? open() : close();
      }}
    >
      <Dialog.Trigger asChild>
        <Button variant="secondary">チャンネルを作る</Button>
      </Dialog.Trigger>
      <Dialog.Overlay />
      <Dialog.Content>
        <Dialog.RightTopClose />
        <Dialog.Title>チャンネルを作る</Dialog.Title>
        <Dialog.Description>
          チャンネルを作って、みんなで話しましょう
        </Dialog.Description>
        <form
          onSubmit={handleSubmit(onSubmit)}
          className={styles.createChannelForm}
        >
          <TextInput
            label="チャンネル名"
            id="channel-name"
            error={errors.name?.message}
            {...register("name")}
          />
          {formError && <p className={styles.formError}>{formError}</p>}
          <Dialog.Footer>
            <Dialog.Close asChild>
              <Button variant="secondary">Cancel</Button>
            </Dialog.Close>
            <Button variant="primary" type="submit">
              OK
            </Button>
          </Dialog.Footer>
        </form>
      </Dialog.Content>
    </Dialog.Root>
  );
};
