import utils from './utils';
import { Either, left, right } from './either';

class ClusterService {
  endpoint: string;

  constructor(endpoint: string) {
    this.endpoint = endpoint;
  }

  public getClusters(): Promise<Cluster[]> {
    return (
      utils
        .restCall(this.endpoint + '/clusters', 'GET')
        .then((response) => {
          if (response.ok) {
            return response.json();
          }
          throw response;
        })
        //.catch(error => console.log('error: ', error))
        .then((clusters) =>
          clusters.map(
            (cluster) =>
              <Cluster>{
                clusterName: cluster.name,
                clusterNamespace: cluster.namespace,
                clusterType: cluster.type,
                clusterStatus: cluster.status,
                clusterConsoleUri: cluster.console,
              }
          )
        ) as Promise<Cluster[]>
    );
  }
}

const clusterService: ClusterService = new ClusterService(utils.endpoint());

export default clusterService;
