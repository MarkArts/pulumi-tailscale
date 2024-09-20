// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.tailscale.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.tailscale.inputs.ContactsAccountArgs;
import com.pulumi.tailscale.inputs.ContactsSecurityArgs;
import com.pulumi.tailscale.inputs.ContactsSupportArgs;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ContactsState extends com.pulumi.resources.ResourceArgs {

    public static final ContactsState Empty = new ContactsState();

    /**
     * Configuration for communications about important changes to your tailnet
     * 
     */
    @Import(name="account")
    private @Nullable Output<ContactsAccountArgs> account;

    /**
     * @return Configuration for communications about important changes to your tailnet
     * 
     */
    public Optional<Output<ContactsAccountArgs>> account() {
        return Optional.ofNullable(this.account);
    }

    /**
     * Configuration for communications about security issues affecting your tailnet
     * 
     */
    @Import(name="security")
    private @Nullable Output<ContactsSecurityArgs> security;

    /**
     * @return Configuration for communications about security issues affecting your tailnet
     * 
     */
    public Optional<Output<ContactsSecurityArgs>> security() {
        return Optional.ofNullable(this.security);
    }

    /**
     * Configuration for communications about misconfigurations in your tailnet
     * 
     */
    @Import(name="support")
    private @Nullable Output<ContactsSupportArgs> support;

    /**
     * @return Configuration for communications about misconfigurations in your tailnet
     * 
     */
    public Optional<Output<ContactsSupportArgs>> support() {
        return Optional.ofNullable(this.support);
    }

    private ContactsState() {}

    private ContactsState(ContactsState $) {
        this.account = $.account;
        this.security = $.security;
        this.support = $.support;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ContactsState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ContactsState $;

        public Builder() {
            $ = new ContactsState();
        }

        public Builder(ContactsState defaults) {
            $ = new ContactsState(Objects.requireNonNull(defaults));
        }

        /**
         * @param account Configuration for communications about important changes to your tailnet
         * 
         * @return builder
         * 
         */
        public Builder account(@Nullable Output<ContactsAccountArgs> account) {
            $.account = account;
            return this;
        }

        /**
         * @param account Configuration for communications about important changes to your tailnet
         * 
         * @return builder
         * 
         */
        public Builder account(ContactsAccountArgs account) {
            return account(Output.of(account));
        }

        /**
         * @param security Configuration for communications about security issues affecting your tailnet
         * 
         * @return builder
         * 
         */
        public Builder security(@Nullable Output<ContactsSecurityArgs> security) {
            $.security = security;
            return this;
        }

        /**
         * @param security Configuration for communications about security issues affecting your tailnet
         * 
         * @return builder
         * 
         */
        public Builder security(ContactsSecurityArgs security) {
            return security(Output.of(security));
        }

        /**
         * @param support Configuration for communications about misconfigurations in your tailnet
         * 
         * @return builder
         * 
         */
        public Builder support(@Nullable Output<ContactsSupportArgs> support) {
            $.support = support;
            return this;
        }

        /**
         * @param support Configuration for communications about misconfigurations in your tailnet
         * 
         * @return builder
         * 
         */
        public Builder support(ContactsSupportArgs support) {
            return support(Output.of(support));
        }

        public ContactsState build() {
            return $;
        }
    }

}