// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.tailscale.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetAclResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The contents of Tailscale ACL as JSON
     * 
     */
    private String json;

    private GetAclResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The contents of Tailscale ACL as JSON
     * 
     */
    public String json() {
        return this.json;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAclResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private String json;
        public Builder() {}
        public Builder(GetAclResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.json = defaults.json;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder json(String json) {
            this.json = Objects.requireNonNull(json);
            return this;
        }
        public GetAclResult build() {
            final var _resultValue = new GetAclResult();
            _resultValue.id = id;
            _resultValue.json = json;
            return _resultValue;
        }
    }
}
