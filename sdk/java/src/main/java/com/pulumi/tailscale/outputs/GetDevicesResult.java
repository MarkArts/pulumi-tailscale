// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.tailscale.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.tailscale.outputs.GetDevicesDevice;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetDevicesResult {
    /**
     * @return The list of devices in the tailnet
     * 
     */
    private List<GetDevicesDevice> devices;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Filters the device list to elements whose name has the provided prefix
     * 
     */
    private @Nullable String namePrefix;

    private GetDevicesResult() {}
    /**
     * @return The list of devices in the tailnet
     * 
     */
    public List<GetDevicesDevice> devices() {
        return this.devices;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Filters the device list to elements whose name has the provided prefix
     * 
     */
    public Optional<String> namePrefix() {
        return Optional.ofNullable(this.namePrefix);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDevicesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetDevicesDevice> devices;
        private String id;
        private @Nullable String namePrefix;
        public Builder() {}
        public Builder(GetDevicesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.devices = defaults.devices;
    	      this.id = defaults.id;
    	      this.namePrefix = defaults.namePrefix;
        }

        @CustomType.Setter
        public Builder devices(List<GetDevicesDevice> devices) {
            this.devices = Objects.requireNonNull(devices);
            return this;
        }
        public Builder devices(GetDevicesDevice... devices) {
            return devices(List.of(devices));
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder namePrefix(@Nullable String namePrefix) {
            this.namePrefix = namePrefix;
            return this;
        }
        public GetDevicesResult build() {
            final var _resultValue = new GetDevicesResult();
            _resultValue.devices = devices;
            _resultValue.id = id;
            _resultValue.namePrefix = namePrefix;
            return _resultValue;
        }
    }
}
