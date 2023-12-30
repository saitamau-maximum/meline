import { Dialog } from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
import { useState } from "react";
import { useCreateChannelForm } from "./use-create-channel-form";
import { TextInput } from "@/components/ui/text-input";
import { styles } from "./create-channel-form.css";
import { IChannelRepository } from "@/repositories/channel";

interface CreateChannelFormProps {
  fetchJoinedChannels: () => Promise<void>;
  channelRepository: IChannelRepository;
}

export const CreateChannelForm = ({
  fetchJoinedChannels,
  channelRepository,
}: CreateChannelFormProps) => {
  const [isOpen, setIsOpen] = useState(false);

  const open = () => setIsOpen(true);
  const close = () => {
    reset();
    setIsOpen(false);
  };

  const onCreated = () => {
    void fetchJoinedChannels();
    close();
  };

  const { reset, register, errors, handleSubmit } = useCreateChannelForm({
    channelRepository,
    onCreated,
  });

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
          onSubmit={(d) => void handleSubmit(d)}
          className={styles.createChannelForm}
        >
          <TextInput
            label="チャンネル名"
            id="channel-name"
            error={errors.name?.message}
            {...register("name")}
          />
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
