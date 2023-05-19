import { singleton } from "tsyringe";
import { EnvProvider, IRawEnvironment } from "../providers/env.provider";
import { NodeEnv } from "../models/enum/node-env.enum";

@singleton()
export class ConfigService {
    private readonly rawEnv: IRawEnvironment;

    /**
     * The port the application should be running on and listening for requests.
     */
    public get AppPort(): number {
        return parseInt(this.rawEnv.APP_PORT);
    }

    public get NodeEnv(): NodeEnv {
        return this.rawEnv.NODE_ENV.toLowerCase() as NodeEnv;
    }

    public get DBHost(): string {
        return this.rawEnv.DB_HOST.toString();
    }

    public get DBPort(): number {
        return parseInt(this.rawEnv.DB_PORT);
    }

    public get DBUsername(): string {
        return this.rawEnv.DB_USERNAME.toString();
    }

    public get DBPassword(): string {
        return this.rawEnv.DB_PASSWORD.toString();
    }

    public get DBName(): string {
        return this.rawEnv.DB_NAME.toString();
    }

    public get DBSync(): boolean {
        return this.rawEnv.DB_SYNC.toLowerCase() === "true";
    }

    public get DBLogging(): boolean {
        return this.rawEnv.DB_LOGGING.toLowerCase() === "true";
    }

    public get RedisHost(): string {
        return this.rawEnv.REDIS_HOST.toString();
    }

    constructor() {
        const envProvider = new EnvProvider();
        this.rawEnv = envProvider.RawEnv;
    }
}