class Utils {
  public isDevMode(): boolean {
    return !process.env.NODE_ENV || process.env.NODE_ENV === 'development';
  }

  /**
   * Decide the endpoint depending if we are in dev mode or production mode
   */
  public endpoint(): string {
    if (this.isDevMode()) {
      return 'http://localhost:9001/rest';
    } else {
      return window.location.origin.toString() + '/rest';
    }
  }

  /**
   * Perform a REST call
   *
   * @param url
   * @param method
   */
  public restCall(url: string, method: string): Promise<Response> {
    return fetch(url, {
      method: method,
    });
  }

  public mapError(err: any, errorMessage?: string): ActionResponse {
    console.error(err);
    if (err instanceof TypeError) {
      return <ActionResponse>{
        message: !errorMessage ? err.message : errorMessage,
        success: false,
      };
    }

    if (err instanceof Response) {
      if (err.status == 401) {
        return <ActionResponse>{
          message: 'Login failed. Check your credentials and try again.',
          success: false,
        };
      }
    }

    return err.text().then(
      (text) =>
        <ActionResponse>{
          message: !errorMessage ? text : errorMessage,
          success: false,
        }
    );
  }
}

const utils: Utils = new Utils();

export default utils;
