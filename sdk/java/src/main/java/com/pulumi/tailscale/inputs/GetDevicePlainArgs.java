// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.tailscale.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetDevicePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetDevicePlainArgs Empty = new GetDevicePlainArgs();

    /**
     * The name of the device
     * 
     */
    @Import(name="name", required=true)
    private String name;

    /**
     * @return The name of the device
     * 
     */
    public String name() {
        return this.name;
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
         * @param name The name of the device
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
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
            if ($.name == null) {
                throw new MissingRequiredPropertyException("GetDevicePlainArgs", "name");
            }
            return $;
        }
    }

}
