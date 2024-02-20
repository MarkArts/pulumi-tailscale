// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.tailscale.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetDevicePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetDevicePlainArgs Empty = new GetDevicePlainArgs();

    /**
     * The short hostname of the device
     * 
     */
    @Import(name="hostname")
    private @Nullable String hostname;

    /**
     * @return The short hostname of the device
     * 
     */
    public Optional<String> hostname() {
        return Optional.ofNullable(this.hostname);
    }

    /**
     * The full name of the device (e.g. `hostname.domain.ts.net`)
     * 
     */
    @Import(name="name")
    private @Nullable String name;

    /**
     * @return The full name of the device (e.g. `hostname.domain.ts.net`)
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * If specified, the provider will make multiple attempts to obtain the data source until the wait_for duration is reached. Retries are made every second so this value should be greater than 1s
     * 
     */
    @Import(name="waitFor")
    private @Nullable String waitFor;

    /**
     * @return If specified, the provider will make multiple attempts to obtain the data source until the wait_for duration is reached. Retries are made every second so this value should be greater than 1s
     * 
     */
    public Optional<String> waitFor() {
        return Optional.ofNullable(this.waitFor);
    }

    private GetDevicePlainArgs() {}

    private GetDevicePlainArgs(GetDevicePlainArgs $) {
        this.hostname = $.hostname;
        this.name = $.name;
        this.waitFor = $.waitFor;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetDevicePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetDevicePlainArgs $;

        public Builder() {
            $ = new GetDevicePlainArgs();
        }

        public Builder(GetDevicePlainArgs defaults) {
            $ = new GetDevicePlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param hostname The short hostname of the device
         * 
         * @return builder
         * 
         */
        public Builder hostname(@Nullable String hostname) {
            $.hostname = hostname;
            return this;
        }

        /**
         * @param name The full name of the device (e.g. `hostname.domain.ts.net`)
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable String name) {
            $.name = name;
            return this;
        }

        /**
         * @param waitFor If specified, the provider will make multiple attempts to obtain the data source until the wait_for duration is reached. Retries are made every second so this value should be greater than 1s
         * 
         * @return builder
         * 
         */
        public Builder waitFor(@Nullable String waitFor) {
            $.waitFor = waitFor;
            return this;
        }

        public GetDevicePlainArgs build() {
            return $;
        }
    }

}
