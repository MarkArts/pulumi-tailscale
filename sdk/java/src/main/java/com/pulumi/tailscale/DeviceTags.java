// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.tailscale;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.tailscale.DeviceTagsArgs;
import com.pulumi.tailscale.Utilities;
import com.pulumi.tailscale.inputs.DeviceTagsState;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * The device_tags resource is used to apply tags to Tailscale devices. See https://tailscale.com/kb/1068/acl-tags/ for more details.
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
 * import com.pulumi.tailscale.TailscaleFunctions;
 * import com.pulumi.tailscale.inputs.GetDeviceArgs;
 * import com.pulumi.tailscale.DeviceTags;
 * import com.pulumi.tailscale.DeviceTagsArgs;
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
 *         final var sampleDevice = TailscaleFunctions.getDevice(GetDeviceArgs.builder()
 *             .name("device.example.com")
 *             .build());
 * 
 *         var sampleTags = new DeviceTags("sampleTags", DeviceTagsArgs.builder()
 *             .deviceId(sampleDevice.applyValue(getDeviceResult -> getDeviceResult.id()))
 *             .tags("room:bedroom")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Device tags can be imported using the device id, e.g.,
 * 
 * ```sh
 * $ pulumi import tailscale:index/deviceTags:DeviceTags sample 123456789
 * ```
 * 
 */
@ResourceType(type="tailscale:index/deviceTags:DeviceTags")
public class DeviceTags extends com.pulumi.resources.CustomResource {
    /**
     * The device to set tags for
     * 
     */
    @Export(name="deviceId", refs={String.class}, tree="[0]")
    private Output<String> deviceId;

    /**
     * @return The device to set tags for
     * 
     */
    public Output<String> deviceId() {
        return this.deviceId;
    }
    /**
     * The tags to apply to the device
     * 
     */
    @Export(name="tags", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> tags;

    /**
     * @return The tags to apply to the device
     * 
     */
    public Output<List<String>> tags() {
        return this.tags;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DeviceTags(java.lang.String name) {
        this(name, DeviceTagsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DeviceTags(java.lang.String name, DeviceTagsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DeviceTags(java.lang.String name, DeviceTagsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("tailscale:index/deviceTags:DeviceTags", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private DeviceTags(java.lang.String name, Output<java.lang.String> id, @Nullable DeviceTagsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("tailscale:index/deviceTags:DeviceTags", name, state, makeResourceOptions(options, id), false);
    }

    private static DeviceTagsArgs makeArgs(DeviceTagsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? DeviceTagsArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static DeviceTags get(java.lang.String name, Output<java.lang.String> id, @Nullable DeviceTagsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DeviceTags(name, id, state, options);
    }
}
