// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.tailscale.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TailnetKeyState extends com.pulumi.resources.ResourceArgs {

    public static final TailnetKeyState Empty = new TailnetKeyState();

    /**
     * The creation timestamp of the key in RFC3339 format
     * 
     */
    @Import(name="createdAt")
    private @Nullable Output<String> createdAt;

    /**
     * @return The creation timestamp of the key in RFC3339 format
     * 
     */
    public Optional<Output<String>> createdAt() {
        return Optional.ofNullable(this.createdAt);
    }

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
     * The expiry timestamp of the key in RFC3339 format
     * 
     */
    @Import(name="expiresAt")
    private @Nullable Output<String> expiresAt;

    /**
     * @return The expiry timestamp of the key in RFC3339 format
     * 
     */
    public Optional<Output<String>> expiresAt() {
        return Optional.ofNullable(this.expiresAt);
    }

    /**
     * The expiry of the key in seconds
     * 
     */
    @Import(name="expiry")
    private @Nullable Output<Integer> expiry;

    /**
     * @return The expiry of the key in seconds
     * 
     */
    public Optional<Output<Integer>> expiry() {
        return Optional.ofNullable(this.expiry);
    }

    /**
     * The authentication key
     * 
     */
    @Import(name="key")
    private @Nullable Output<String> key;

    /**
     * @return The authentication key
     * 
     */
    public Optional<Output<String>> key() {
        return Optional.ofNullable(this.key);
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

    private TailnetKeyState() {}

    private TailnetKeyState(TailnetKeyState $) {
        this.createdAt = $.createdAt;
        this.ephemeral = $.ephemeral;
        this.expiresAt = $.expiresAt;
        this.expiry = $.expiry;
        this.key = $.key;
        this.preauthorized = $.preauthorized;
        this.reusable = $.reusable;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TailnetKeyState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TailnetKeyState $;

        public Builder() {
            $ = new TailnetKeyState();
        }

        public Builder(TailnetKeyState defaults) {
            $ = new TailnetKeyState(Objects.requireNonNull(defaults));
        }

        /**
         * @param createdAt The creation timestamp of the key in RFC3339 format
         * 
         * @return builder
         * 
         */
        public Builder createdAt(@Nullable Output<String> createdAt) {
            $.createdAt = createdAt;
            return this;
        }

        /**
         * @param createdAt The creation timestamp of the key in RFC3339 format
         * 
         * @return builder
         * 
         */
        public Builder createdAt(String createdAt) {
            return createdAt(Output.of(createdAt));
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
         * @param expiresAt The expiry timestamp of the key in RFC3339 format
         * 
         * @return builder
         * 
         */
        public Builder expiresAt(@Nullable Output<String> expiresAt) {
            $.expiresAt = expiresAt;
            return this;
        }

        /**
         * @param expiresAt The expiry timestamp of the key in RFC3339 format
         * 
         * @return builder
         * 
         */
        public Builder expiresAt(String expiresAt) {
            return expiresAt(Output.of(expiresAt));
        }

        /**
         * @param expiry The expiry of the key in seconds
         * 
         * @return builder
         * 
         */
        public Builder expiry(@Nullable Output<Integer> expiry) {
            $.expiry = expiry;
            return this;
        }

        /**
         * @param expiry The expiry of the key in seconds
         * 
         * @return builder
         * 
         */
        public Builder expiry(Integer expiry) {
            return expiry(Output.of(expiry));
        }

        /**
         * @param key The authentication key
         * 
         * @return builder
         * 
         */
        public Builder key(@Nullable Output<String> key) {
            $.key = key;
            return this;
        }

        /**
         * @param key The authentication key
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            return key(Output.of(key));
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

        public TailnetKeyState build() {
            return $;
        }
    }

}
