// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.tailscale.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetDeviceResult {
    /**
     * @return The list of device&#39;s IPs
     * 
     */
    private List<String> addresses;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The name of the device
     * 
     */
    private String name;
    /**
     * @return The tags applied to the device
     * 
     */
    private List<String> tags;
    /**
     * @return The user associated with the device
     * 
     */
    private String user;
    /**
     * @return If specified, the provider will make multiple attempts to obtain the data source until the wait_for duration is reached. Retries are made every second so this value should be greater than 1s
     * 
     */
    private @Nullable String waitFor;

    private GetDeviceResult() {}
    /**
     * @return The list of device&#39;s IPs
     * 
     */
    public List<String> addresses() {
        return this.addresses;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The name of the device
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return The tags applied to the device
     * 
     */
    public List<String> tags() {
        return this.tags;
    }
    /**
     * @return The user associated with the device
     * 
     */
    public String user() {
        return this.user;
    }
    /**
     * @return If specified, the provider will make multiple attempts to obtain the data source until the wait_for duration is reached. Retries are made every second so this value should be greater than 1s
     * 
     */
    public Optional<String> waitFor() {
        return Optional.ofNullable(this.waitFor);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDeviceResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> addresses;
        private String id;
        private String name;
        private List<String> tags;
        private String user;
        private @Nullable String waitFor;
        public Builder() {}
        public Builder(GetDeviceResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.addresses = defaults.addresses;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.tags = defaults.tags;
    	      this.user = defaults.user;
    	      this.waitFor = defaults.waitFor;
        }

        @CustomType.Setter
        public Builder addresses(List<String> addresses) {
            this.addresses = Objects.requireNonNull(addresses);
            return this;
        }
        public Builder addresses(String... addresses) {
            return addresses(List.of(addresses));
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder tags(List<String> tags) {
            this.tags = Objects.requireNonNull(tags);
            return this;
        }
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }
        @CustomType.Setter
        public Builder user(String user) {
            this.user = Objects.requireNonNull(user);
            return this;
        }
        @CustomType.Setter
        public Builder waitFor(@Nullable String waitFor) {
            this.waitFor = waitFor;
            return this;
        }
        public GetDeviceResult build() {
            final var o = new GetDeviceResult();
            o.addresses = addresses;
            o.id = id;
            o.name = name;
            o.tags = tags;
            o.user = user;
            o.waitFor = waitFor;
            return o;
        }
    }
}
