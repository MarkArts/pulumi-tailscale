// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.tailscale;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.tailscale.TailnetKeyArgs;
import com.pulumi.tailscale.Utilities;
import com.pulumi.tailscale.inputs.TailnetKeyState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The tailnet_key resource allows you to create pre-authentication keys that can register new nodes without needing to sign in via a web browser. See https://tailscale.com/kb/1085/auth-keys for more information
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.tailscale.TailnetKey;
 * import com.pulumi.tailscale.TailnetKeyArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var sampleKey = new TailnetKey(&#34;sampleKey&#34;, TailnetKeyArgs.builder()        
 *             .ephemeral(false)
 *             .expiry(3600)
 *             .preauthorized(true)
 *             .reusable(true)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="tailscale:index/tailnetKey:TailnetKey")
public class TailnetKey extends com.pulumi.resources.CustomResource {
    /**
     * The creation timestamp of the key in RFC3339 format
     * 
     */
    @Export(name="createdAt", type=String.class, parameters={})
    private Output<String> createdAt;

    /**
     * @return The creation timestamp of the key in RFC3339 format
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * Indicates if the key is ephemeral.
     * 
     */
    @Export(name="ephemeral", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> ephemeral;

    /**
     * @return Indicates if the key is ephemeral.
     * 
     */
    public Output<Optional<Boolean>> ephemeral() {
        return Codegen.optional(this.ephemeral);
    }
    /**
     * The expiry timestamp of the key in RFC3339 format
     * 
     */
    @Export(name="expiresAt", type=String.class, parameters={})
    private Output<String> expiresAt;

    /**
     * @return The expiry timestamp of the key in RFC3339 format
     * 
     */
    public Output<String> expiresAt() {
        return this.expiresAt;
    }
    /**
     * The expiry of the key in seconds
     * 
     */
    @Export(name="expiry", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> expiry;

    /**
     * @return The expiry of the key in seconds
     * 
     */
    public Output<Optional<Integer>> expiry() {
        return Codegen.optional(this.expiry);
    }
    /**
     * The authentication key
     * 
     */
    @Export(name="key", type=String.class, parameters={})
    private Output<String> key;

    /**
     * @return The authentication key
     * 
     */
    public Output<String> key() {
        return this.key;
    }
    /**
     * Determines whether or not the machines authenticated by the key will be authorized for the tailnet by default.
     * 
     */
    @Export(name="preauthorized", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> preauthorized;

    /**
     * @return Determines whether or not the machines authenticated by the key will be authorized for the tailnet by default.
     * 
     */
    public Output<Optional<Boolean>> preauthorized() {
        return Codegen.optional(this.preauthorized);
    }
    /**
     * Indicates if the key is reusable or single-use.
     * 
     */
    @Export(name="reusable", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> reusable;

    /**
     * @return Indicates if the key is reusable or single-use.
     * 
     */
    public Output<Optional<Boolean>> reusable() {
        return Codegen.optional(this.reusable);
    }
    /**
     * List of tags to apply to the machines authenticated by the key.
     * 
     */
    @Export(name="tags", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> tags;

    /**
     * @return List of tags to apply to the machines authenticated by the key.
     * 
     */
    public Output<Optional<List<String>>> tags() {
        return Codegen.optional(this.tags);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public TailnetKey(String name) {
        this(name, TailnetKeyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public TailnetKey(String name, @Nullable TailnetKeyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public TailnetKey(String name, @Nullable TailnetKeyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("tailscale:index/tailnetKey:TailnetKey", name, args == null ? TailnetKeyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private TailnetKey(String name, Output<String> id, @Nullable TailnetKeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("tailscale:index/tailnetKey:TailnetKey", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "key"
            ))
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static TailnetKey get(String name, Output<String> id, @Nullable TailnetKeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new TailnetKey(name, id, state, options);
    }
}
