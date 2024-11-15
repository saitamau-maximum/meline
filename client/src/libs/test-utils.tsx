import {
  RepositoryContext,
  RepositoryContextProps,
} from "@/providers/repository";
import { DefaultRepositories } from "@/repositories";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { render, renderHook, RenderOptions } from "@testing-library/react";

export const generateWrapper = (
  repositories: Partial<RepositoryContextProps>
) => {
  const initialRepositories = {
    ...DefaultRepositories,
    ...repositories,
  };
  const queryClient = new QueryClient();

  return ({ children }: { children: React.ReactNode }) => {
    return (
      <RepositoryContext.Provider value={initialRepositories}>
        <QueryClientProvider client={queryClient}>
          {children}
        </QueryClientProvider>
      </RepositoryContext.Provider>
    );
  };
};

interface CustomRenderOptions {
  repositories?: Partial<RepositoryContextProps>;
}

const customRender = <
  Q extends typeof import("@testing-library/dom").queries,
  Container extends HTMLElement | null = HTMLElement,
  BaseElement extends HTMLElement | null = HTMLElement
>(
  ui: React.ReactElement,
  options: RenderOptions<Q, Container, BaseElement> & CustomRenderOptions = {}
) => {
  const { repositories, ...renderOptions } = options;
  return render(ui, {
    wrapper: generateWrapper(repositories || {}),
    ...renderOptions,
  });
};

const customRenderHook = <
  Result,
  Props,
  Q extends typeof import("@testing-library/dom").queries,
  Container extends HTMLElement | null = HTMLElement,
  BaseElement extends HTMLElement | null = HTMLElement
>(
  render: (initialProps: Props) => Result,
  options: RenderOptions<Q, Container, BaseElement> & CustomRenderOptions = {}
) => {
  const { repositories, ...renderOptions } = options;
  return renderHook(render, {
    wrapper: generateWrapper(repositories || {}),
    ...renderOptions,
  });
};

export * from "@testing-library/react";

export { customRender as render, customRenderHook as renderHook };
