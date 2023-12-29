import { Meta } from "@storybook/react";
import { LoadingOverlay } from "..";
import { LoadingOverlayContext } from "@/providers/loading-overlay";

const meta = {
  title: "Components/LoadingOverlay",
} satisfies Meta;

export default meta;

export const Overview = () => (
  <LoadingOverlayContext.Provider
    value={{ isLoading: true, setIsLoading: () => {} }}
  >
    <LoadingOverlay />
  </LoadingOverlayContext.Provider>
);
