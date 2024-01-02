// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.tailscale.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;


public final class Get4Via6Args extends com.pulumi.resources.InvokeArgs {

    public static final Get4Via6Args Empty = new Get4Via6Args();

    /**
     * The IPv4 CIDR to map
     * 
     */
    @Import(name="cidr", required=true)
    private Output<String> cidr;

    /**
     * @return The IPv4 CIDR to map
     * 
     */
    public Output<String> cidr() {
        return this.cidr;
    }

    /**
     * Site ID (between 0 and 255)
     * 
     */
    @Import(name="site", required=true)
    private Output<Integer> site;

    /**
     * @return Site ID (between 0 and 255)
     * 
     */
    public Output<Integer> site() {
        return this.site;
    }

    private Get4Via6Args() {}

    private Get4Via6Args(Get4Via6Args $) {
        this.cidr = $.cidr;
        this.site = $.site;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(Get4Via6Args defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private Get4Via6Args $;

        public Builder() {
            $ = new Get4Via6Args();
        }

        public Builder(Get4Via6Args defaults) {
            $ = new Get4Via6Args(Objects.requireNonNull(defaults));
        }

        /**
         * @param cidr The IPv4 CIDR to map
         * 
         * @return builder
         * 
         */
        public Builder cidr(Output<String> cidr) {
            $.cidr = cidr;
            return this;
        }

        /**
         * @param cidr The IPv4 CIDR to map
         * 
         * @return builder
         * 
         */
        public Builder cidr(String cidr) {
            return cidr(Output.of(cidr));
        }

        /**
         * @param site Site ID (between 0 and 255)
         * 
         * @return builder
         * 
         */
        public Builder site(Output<Integer> site) {
            $.site = site;
            return this;
        }

        /**
         * @param site Site ID (between 0 and 255)
         * 
         * @return builder
         * 
         */
        public Builder site(Integer site) {
            return site(Output.of(site));
        }

        public Get4Via6Args build() {
            if ($.cidr == null) {
                throw new MissingRequiredPropertyException("Get4Via6Args", "cidr");
            }
            if ($.site == null) {
                throw new MissingRequiredPropertyException("Get4Via6Args", "site");
            }
            return $;
        }
    }

}
