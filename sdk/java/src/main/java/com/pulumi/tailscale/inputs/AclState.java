// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.tailscale.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AclState extends com.pulumi.resources.ResourceArgs {

    public static final AclState Empty = new AclState();

    /**
     * The policy that defines which devices and users are allowed to connect in your network. Can be either a JSON or a HuJSON string.
     * 
     */
    @Import(name="acl")
    private @Nullable Output<String> acl;

    /**
     * @return The policy that defines which devices and users are allowed to connect in your network. Can be either a JSON or a HuJSON string.
     * 
     */
    public Optional<Output<String>> acl() {
        return Optional.ofNullable(this.acl);
    }

    /**
     * If true, will skip requirement to import acl before allowing changes. Be careful, can cause ACL to be overwritten
     * 
     */
    @Import(name="overwriteExistingContent")
    private @Nullable Output<Boolean> overwriteExistingContent;

    /**
     * @return If true, will skip requirement to import acl before allowing changes. Be careful, can cause ACL to be overwritten
     * 
     */
    public Optional<Output<Boolean>> overwriteExistingContent() {
        return Optional.ofNullable(this.overwriteExistingContent);
    }

    private AclState() {}

    private AclState(AclState $) {
        this.acl = $.acl;
        this.overwriteExistingContent = $.overwriteExistingContent;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AclState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AclState $;

        public Builder() {
            $ = new AclState();
        }

        public Builder(AclState defaults) {
            $ = new AclState(Objects.requireNonNull(defaults));
        }

        /**
         * @param acl The policy that defines which devices and users are allowed to connect in your network. Can be either a JSON or a HuJSON string.
         * 
         * @return builder
         * 
         */
        public Builder acl(@Nullable Output<String> acl) {
            $.acl = acl;
            return this;
        }

        /**
         * @param acl The policy that defines which devices and users are allowed to connect in your network. Can be either a JSON or a HuJSON string.
         * 
         * @return builder
         * 
         */
        public Builder acl(String acl) {
            return acl(Output.of(acl));
        }

        /**
         * @param overwriteExistingContent If true, will skip requirement to import acl before allowing changes. Be careful, can cause ACL to be overwritten
         * 
         * @return builder
         * 
         */
        public Builder overwriteExistingContent(@Nullable Output<Boolean> overwriteExistingContent) {
            $.overwriteExistingContent = overwriteExistingContent;
            return this;
        }

        /**
         * @param overwriteExistingContent If true, will skip requirement to import acl before allowing changes. Be careful, can cause ACL to be overwritten
         * 
         * @return builder
         * 
         */
        public Builder overwriteExistingContent(Boolean overwriteExistingContent) {
            return overwriteExistingContent(Output.of(overwriteExistingContent));
        }

        public AclState build() {
            return $;
        }
    }

}
