import { Meta } from "@storybook/react";
import { Theme } from "@/styles/__stories__/Theme";
import { LoadingOverlay } from "..";
import { LoadingOverlayContext } from "@/providers/loading-overlay";

const meta = {
  title: "Components/LoadingOverlay",
} satisfies Meta;

export default meta;

export const Overview = () => (
  <Theme.Light>
    <LoadingOverlayContext.Provider
      value={{ isLoading: true, setIsLoading: () => {} }}
    >
      <LoadingOverlay />
    </LoadingOverlayContext.Provider>
  </Theme.Light>
);
