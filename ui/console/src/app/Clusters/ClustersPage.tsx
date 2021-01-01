import React from "react";
import { Link, useHistory } from "react-router-dom";
import { useTranslation } from "react-i18next";
import {
  AlertVariant,
  Badge,
  Button,
  PageSection,
} from "@patternfly/react-core";

import { ViewHeader } from "../components/view-header/ViewHeader";
import { ConsoleDataTable } from "../components/table-toolbar/ConsoleDataTable";
import { IFormatter, IFormatterValueType } from "@patternfly/react-table";
import { useAlerts } from "../components/alert/Alerts";
import { ExternalLink } from "../components/external-link/ExternalLink";
import clusterService from "../../services/clusterService";

export const ClustersPage = () => {
  const { t } = useTranslation("clients");
  const { addAlert } = useAlerts();
  const history = useHistory();

  const loader = async (first?: number, max?: number, search?: string) => {
    try {
      return await clusterService.getClusters()
    } catch (ex) {
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
        titleKey="clients:clientList"
        subKey="clients:clientsExplain"
      />
      <PageSection variant="light">
        <ConsoleDataTable
          loader={loader}
          isPaginated
          ariaLabelKey="clients:clientList"
          searchPlaceholderKey="clients:searchForClient"
          toolbarItem={
            <>
              <Button onClick={() => history.push("/add-client")}>
                {t("createClient")}
              </Button>
              <Button
                onClick={() => history.push("/import-client")}
                variant="link"
              >
                {t("importClient")}
              </Button>
            </>
          }
          columns={[
            {
              name: "clusterName",
              displayKey: "clients:clientID",
            },
            { name: "clusterNamespace", displayKey: "clients:type" },
            /*
            {
              name: "description",
              displayKey: "clients:description",
              cellFormatters: [emptyFormatter()],
            },
            {
              name: "baseUrl",
              displayKey: "clients:homeURL",
              cellFormatters: [externalLink(), emptyFormatter()],

              cellRenderer: (client) => {
                if (client.rootUrl) {
                  if (
                    !client.rootUrl.startsWith("http") ||
                    client.rootUrl.indexOf("$") !== -1
                  ) {
                    client.rootUrl =
                      client.rootUrl
                        .replace("${authBaseUrl}", baseUrl)
                        .replace("${authAdminUrl}", baseUrl) +
                      (client.baseUrl ? client.baseUrl.substr(1) : "");
                  }
                }
                return client.rootUrl;
              },

            },
             */
          ]}
        />
      </PageSection>
    </>
  );
};
