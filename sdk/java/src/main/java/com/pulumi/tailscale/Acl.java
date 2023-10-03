// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.tailscale;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.tailscale.AclArgs;
import com.pulumi.tailscale.Utilities;
import com.pulumi.tailscale.inputs.AclState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * The acl resource allows you to configure a Tailscale ACL. See https://tailscale.com/kb/1018/acls for more information. Note that this resource will completely overwrite existing ACL contents for a given tailnet.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.tailscale.Acl;
 * import com.pulumi.tailscale.AclArgs;
 * import static com.pulumi.codegen.internal.Serialization.*;
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
 *         var sampleAcl = new Acl(&#34;sampleAcl&#34;, AclArgs.builder()        
 *             .acl(serializeJson(
 *                 jsonObject(
 *                     jsonProperty(&#34;acls&#34;, jsonArray(jsonObject(
 *                         jsonProperty(&#34;action&#34;, &#34;accept&#34;),
 *                         jsonProperty(&#34;users&#34;, jsonArray(&#34;*&#34;)),
 *                         jsonProperty(&#34;ports&#34;, jsonArray(&#34;*:*&#34;))
 *                     )))
 *                 )))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="tailscale:index/acl:Acl")
public class Acl extends com.pulumi.resources.CustomResource {
    /**
     * The JSON-based policy that defines which devices and users are allowed to connect in your network
     * 
     */
    @Export(name="acl", refs={String.class}, tree="[0]")
    private Output<String> acl;

    /**
     * @return The JSON-based policy that defines which devices and users are allowed to connect in your network
     * 
     */
    public Output<String> acl() {
        return this.acl;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Acl(String name) {
        this(name, AclArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Acl(String name, AclArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Acl(String name, AclArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("tailscale:index/acl:Acl", name, args == null ? AclArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Acl(String name, Output<String> id, @Nullable AclState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("tailscale:index/acl:Acl", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
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
    public static Acl get(String name, Output<String> id, @Nullable AclState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Acl(name, id, state, options);
    }
}
