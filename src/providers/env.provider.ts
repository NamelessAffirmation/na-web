/**
 * Raw environment variables that we expect to be supplied from the host.
 */
export interface IRawEnvironment {
    APP_PORT: string;
    NODE_ENV: string;
    DB_HOST: string;
    DB_PORT: string;
    DB_USERNAME: string;
    DB_PASSWORD: string;
    DB_NAME: string;
    DB_SYNC: string;
    DB_LOGGING: string;
    REDIS_HOST: string;
}

export class EnvProvider {
    /**
     * When we load the host env variables--this is the raw representation of that data before we verify or type cast it.
     * @private
     */
    private readonly rawEnv: IRawEnvironment;

    public get RawEnv(): IRawEnvironment {
        return this.rawEnv;
    }

    constructor() {
        this.rawEnv = Object.assign({} as IRawEnvironment, process.env);
        console.info("######################");
        console.info("VALIDATING ENV VARIABLES...");
        const errorMessage = this.validateEnv();
        if (errorMessage) {
            console.info(errorMessage);
            throw new Error(
                "Application Failed Env Validation. Please check the host env and correct any issues!"
            );
        } else {
            console.info("Done...");
            console.info("######################");
        }
    }

    /**
     * Validate all the supplied env variables before we allow them to be used by the config service.
     * Returns an ERROR string--if the return is null then we are good to go!
     * If an error string is returned, we will panic and close the application which will halt the init sequence.
     * @private
     */
    private validateEnv(): string | null {
        let errorMessage: string = "";

        if (!this.rawEnv.APP_PORT) {
            errorMessage += "[APP_PORT] is required!\n\n";
        }

        if (!this.rawEnv.NODE_ENV) {
            errorMessage +=
                "[NODE_ENV] is required! Options: ['development','production','testing','migrations']\n\n";
        }

        if (!this.rawEnv.DB_HOST) {
            errorMessage += "[DB_HOST] is required!\n\n";
        }

        if (!this.rawEnv.DB_PORT) {
            errorMessage += "[DB_PORT] is required!\n\n";
        }

        if (!this.rawEnv.DB_USERNAME) {
            errorMessage += "[DB_USERNAME] is required!\n\n";
        }

        if (!this.rawEnv.DB_PASSWORD) {
            errorMessage += "[DB_PASSWORD] is required!\n\n";
        }

        if (!this.rawEnv.DB_NAME) {
            errorMessage += "[DB_NAME] is required!\n\n";
        }

        if (!this.rawEnv.DB_SYNC) {
            errorMessage += "[DB_SYNC] is required!\n\n";
        }

        if (!this.rawEnv.DB_LOGGING) {
            errorMessage += "[DB_LOGGING] is required!\n\n";
        }

        if (!this.rawEnv.REDIS_HOST) {
            errorMessage += "[REDIS_HOST] is required!\n\n";
        }

        return errorMessage ? errorMessage : null;
    }
}