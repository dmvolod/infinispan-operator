import * as React from 'react';
import {
  PageSection,
  Title,
} from '@patternfly/react-core';

export interface ISupportProps {
  sampleProp?: string;
}

const BackupRestore: React.FunctionComponent<ISupportProps> = () => (
    <PageSection>
      <Title headingLevel="h1" size="lg">
        Backup/Restore Settings Page Title
      </Title>
    </PageSection>
  )

export { BackupRestore };
