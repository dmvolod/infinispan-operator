import React, {
  FormEvent,
  Fragment,
  MouseEventHandler,
  ReactNode,
} from "react";
import {
  Toolbar,
  ToolbarContent,
  ToolbarItem,
  InputGroup,
  TextInput,
  Button,
  ButtonVariant,
} from "@patternfly/react-core";
import { SearchIcon } from "@patternfly/react-icons";

type TableToolbarProps = {
  toolbarItem?: ReactNode;
  toolbarItemFooter?: ReactNode;
  children: ReactNode;
  searchTypeComponent?: ReactNode;
  inputGroupName?: string;
  inputGroupPlaceholder?: string;
  inputGroupOnChange?: (
    newInput: string,
    event: FormEvent<HTMLInputElement>
  ) => void;
  inputGroupOnClick?: MouseEventHandler;
};

export const TableToolbar = ({
  toolbarItem,
  toolbarItemFooter,
  children,
  searchTypeComponent,
  inputGroupName,
  inputGroupPlaceholder,
  inputGroupOnChange,
  inputGroupOnClick,
}: TableToolbarProps) => {
  return (
    <>
      <Toolbar>
        <ToolbarContent>
          <Fragment>
            {inputGroupName && (
              <ToolbarItem>
                <InputGroup>
                  {searchTypeComponent}
                  <TextInput
                    name={inputGroupName}
                    id={inputGroupName}
                    type="search"
                    aria-label={"search"}
                    placeholder={inputGroupPlaceholder}
                    onChange={inputGroupOnChange}
                  />
                  <Button
                    variant={ButtonVariant.control}
                    aria-label={"search"}
                    onClick={inputGroupOnClick}
                  >
                    <SearchIcon />
                  </Button>
                </InputGroup>
              </ToolbarItem>
            )}
          </Fragment>
          {toolbarItem}
        </ToolbarContent>
      </Toolbar>
      {children}
      <Toolbar>{toolbarItemFooter}</Toolbar>
    </>
  );
};
