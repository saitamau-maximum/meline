interface ChannelTopPageProps {
  user: {
    name: string;
    imageURL: string;
  };
}

export const ChannelTopPageTemplate = ({ user }: ChannelTopPageProps) => {
  return <div>{JSON.stringify(user)}</div>;
};
