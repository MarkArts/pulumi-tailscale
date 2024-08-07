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
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
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
 *         var sampleKey = new TailnetKey("sampleKey", TailnetKeyArgs.builder()
 *             .reusable(true)
 *             .ephemeral(false)
 *             .preauthorized(true)
 *             .expiry(3600)
 *             .description("Sample key")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="tailscale:index/tailnetKey:TailnetKey")
public class TailnetKey extends com.pulumi.resources.CustomResource {
    /**
     * The creation timestamp of the key in RFC3339 format
     * 
     */
    @Export(name="createdAt", refs={String.class}, tree="[0]")
    private Output<String> createdAt;

    /**
     * @return The creation timestamp of the key in RFC3339 format
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * A description of the key consisting of alphanumeric characters. Defaults to `&#34;&#34;`.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A description of the key consisting of alphanumeric characters. Defaults to `&#34;&#34;`.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Indicates if the key is ephemeral. Defaults to `false`.
     * 
     */
    @Export(name="ephemeral", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> ephemeral;

    /**
     * @return Indicates if the key is ephemeral. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> ephemeral() {
        return Codegen.optional(this.ephemeral);
    }
    /**
     * The expiry timestamp of the key in RFC3339 format
     * 
     */
    @Export(name="expiresAt", refs={String.class}, tree="[0]")
    private Output<String> expiresAt;

    /**
     * @return The expiry timestamp of the key in RFC3339 format
     * 
     */
    public Output<String> expiresAt() {
        return this.expiresAt;
    }
    /**
     * The expiry of the key in seconds. Defaults to `7776000` (90 days).
     * 
     */
    @Export(name="expiry", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> expiry;

    /**
     * @return The expiry of the key in seconds. Defaults to `7776000` (90 days).
     * 
     */
    public Output<Optional<Integer>> expiry() {
        return Codegen.optional(this.expiry);
    }
    /**
     * Indicates whether the key is invalid (e.g. expired, revoked or has been deleted).
     * 
     */
    @Export(name="invalid", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> invalid;

    /**
     * @return Indicates whether the key is invalid (e.g. expired, revoked or has been deleted).
     * 
     */
    public Output<Boolean> invalid() {
        return this.invalid;
    }
    /**
     * The authentication key
     * 
     */
    @Export(name="key", refs={String.class}, tree="[0]")
    private Output<String> key;

    /**
     * @return The authentication key
     * 
     */
    public Output<String> key() {
        return this.key;
    }
    /**
     * Determines whether or not the machines authenticated by the key will be authorized for the tailnet by default. Defaults to `false`.
     * 
     */
    @Export(name="preauthorized", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> preauthorized;

    /**
     * @return Determines whether or not the machines authenticated by the key will be authorized for the tailnet by default. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> preauthorized() {
        return Codegen.optional(this.preauthorized);
    }
    /**
     * Determines whether the key should be created again if it becomes invalid. By default, reusable keys will be recreated, but single-use keys will not. Possible values: &#39;always&#39;, &#39;never&#39;.
     * 
     */
    @Export(name="recreateIfInvalid", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> recreateIfInvalid;

    /**
     * @return Determines whether the key should be created again if it becomes invalid. By default, reusable keys will be recreated, but single-use keys will not. Possible values: &#39;always&#39;, &#39;never&#39;.
     * 
     */
    public Output<Optional<String>> recreateIfInvalid() {
        return Codegen.optional(this.recreateIfInvalid);
    }
    /**
     * Indicates if the key is reusable or single-use. Defaults to `false`.
     * 
     */
    @Export(name="reusable", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> reusable;

    /**
     * @return Indicates if the key is reusable or single-use. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> reusable() {
        return Codegen.optional(this.reusable);
    }
    /**
     * List of tags to apply to the machines authenticated by the key.
     * 
     */
    @Export(name="tags", refs={List.class,String.class}, tree="[0,1]")
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
        super("tailscale:index/tailnetKey:TailnetKey", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()));
    }

    private TailnetKey(String name, Output<String> id, @Nullable TailnetKeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("tailscale:index/tailnetKey:TailnetKey", name, state, makeResourceOptions(options, id));
    }

    private static TailnetKeyArgs makeArgs(@Nullable TailnetKeyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? TailnetKeyArgs.Empty : args;
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
