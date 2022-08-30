// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.tailscale;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TailnetKeyArgs extends com.pulumi.resources.ResourceArgs {

    public static final TailnetKeyArgs Empty = new TailnetKeyArgs();

    /**
     * Indicates if the key is ephemeral.
     * 
     */
    @Import(name="ephemeral")
    private @Nullable Output<Boolean> ephemeral;

    /**
     * @return Indicates if the key is ephemeral.
     * 
     */
    public Optional<Output<Boolean>> ephemeral() {
        return Optional.ofNullable(this.ephemeral);
    }

    /**
     * Determines whether or not the machines authenticated by the key will be authorized for the tailnet by default.
     * 
     */
    @Import(name="preauthorized")
    private @Nullable Output<Boolean> preauthorized;

    /**
     * @return Determines whether or not the machines authenticated by the key will be authorized for the tailnet by default.
     * 
     */
    public Optional<Output<Boolean>> preauthorized() {
        return Optional.ofNullable(this.preauthorized);
    }

    /**
     * Indicates if the key is reusable or single-use.
     * 
     */
    @Import(name="reusable")
    private @Nullable Output<Boolean> reusable;

    /**
     * @return Indicates if the key is reusable or single-use.
     * 
     */
    public Optional<Output<Boolean>> reusable() {
        return Optional.ofNullable(this.reusable);
    }

    /**
     * List of tags to apply to the machines authenticated by the key.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return List of tags to apply to the machines authenticated by the key.
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private TailnetKeyArgs() {}

    private TailnetKeyArgs(TailnetKeyArgs $) {
        this.ephemeral = $.ephemeral;
        this.preauthorized = $.preauthorized;
        this.reusable = $.reusable;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TailnetKeyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TailnetKeyArgs $;

        public Builder() {
            $ = new TailnetKeyArgs();
        }

        public Builder(TailnetKeyArgs defaults) {
            $ = new TailnetKeyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ephemeral Indicates if the key is ephemeral.
         * 
         * @return builder
         * 
         */
        public Builder ephemeral(@Nullable Output<Boolean> ephemeral) {
            $.ephemeral = ephemeral;
            return this;
        }

        /**
         * @param ephemeral Indicates if the key is ephemeral.
         * 
         * @return builder
         * 
         */
        public Builder ephemeral(Boolean ephemeral) {
            return ephemeral(Output.of(ephemeral));
        }

        /**
         * @param preauthorized Determines whether or not the machines authenticated by the key will be authorized for the tailnet by default.
         * 
         * @return builder
         * 
         */
        public Builder preauthorized(@Nullable Output<Boolean> preauthorized) {
            $.preauthorized = preauthorized;
            return this;
        }

        /**
         * @param preauthorized Determines whether or not the machines authenticated by the key will be authorized for the tailnet by default.
         * 
         * @return builder
         * 
         */
        public Builder preauthorized(Boolean preauthorized) {
            return preauthorized(Output.of(preauthorized));
        }

        /**
         * @param reusable Indicates if the key is reusable or single-use.
         * 
         * @return builder
         * 
         */
        public Builder reusable(@Nullable Output<Boolean> reusable) {
            $.reusable = reusable;
            return this;
        }

        /**
         * @param reusable Indicates if the key is reusable or single-use.
         * 
         * @return builder
         * 
         */
        public Builder reusable(Boolean reusable) {
            return reusable(Output.of(reusable));
        }

        /**
         * @param tags List of tags to apply to the machines authenticated by the key.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags List of tags to apply to the machines authenticated by the key.
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags List of tags to apply to the machines authenticated by the key.
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        public TailnetKeyArgs build() {
            return $;
        }
    }

}
