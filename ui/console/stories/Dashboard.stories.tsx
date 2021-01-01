import * as React from 'react';
import { storiesOf } from '@storybook/react';
import { withInfo } from '@storybook/addon-info';
import { Clusters } from '@app/Clusters/Dashboard';

const stories = storiesOf('Components', module);
stories.addDecorator(withInfo);
stories.add(
  'Dashboard',
  () => <Clusters />,
  { info: { inline: true } }
);
