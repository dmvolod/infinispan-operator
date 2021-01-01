import * as React from 'react';
import { storiesOf } from '@storybook/react';
import { withInfo } from '@storybook/addon-info';
import { BackupRestore } from '@app/BackupRestore/Support';

const stories = storiesOf('Components', module);
stories.addDecorator(withInfo);
stories.add(
  'Support',
  () => <BackupRestore />,
  { info: { inline: true } }
);
