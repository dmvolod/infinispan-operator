import React from "react";
import { useTranslation } from "react-i18next";
import {
  PageSection,
} from "@patternfly/react-core";

import { ViewHeader } from "../components/view-header/ViewHeader";
import { ConsoleDataTable } from "../components/table-toolbar/ConsoleDataTable";
import { IFormatter, IFormatterValueType } from "@patternfly/react-table";
import { ExternalLink } from "../components/external-link/ExternalLink";
import clusterService from "../../services/clusterService";

export const ClustersPage = () => {
  const { t } = useTranslation("clients");

  const loader = async (first?: number, max?: number, search?: string) => {
    try {
      return await clusterService.getClusters()
    } catch (error) {
      console.log('error: ', error)
      return [];
    }
  };

  const emptyFormatter = (): IFormatter => (data?: IFormatterValueType) => {
    return data ? data : "—";
  };

  const externalLink = (): IFormatter => (data?: IFormatterValueType) => {
    return (data ? (
      <ExternalLink href={data.toString()} />
    ) : undefined) as object;
  };

  return (
    <>
      <ViewHeader
        titleKey="clusters:clusterList"
        subKey="clusters:clustersExplain"
      />
      <PageSection variant="light">
        <ConsoleDataTable
          loader={loader}
          isPaginated
          ariaLabelKey="clusters:clusterList"
          searchPlaceholderKey="clusters:searchForCluster"
          toolbarItem={
            <>
            </>
          }
          columns={[
            {
              name: "clusterName",
              displayKey: "clusters:clusterName",
            },
            { name: "clusterNamespace",
              displayKey: "clusters:clusterNamespace"
            },
            {
              name: "clusterType",
              displayKey: "clusters:clusterType",
            },
            {
              name: "clusterStatus",
              displayKey: "clusters:clusterStatus",
            },
            {
              name: "clusterConsoleUri",
              displayKey: "clusters:clusterConsoleUri",
              cellFormatters: [externalLink(), emptyFormatter()],
            },
          ]}
        />
      </PageSection>
    </>
  );
};
