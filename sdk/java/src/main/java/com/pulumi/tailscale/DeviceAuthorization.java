// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.tailscale;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.tailscale.DeviceAuthorizationArgs;
import com.pulumi.tailscale.Utilities;
import com.pulumi.tailscale.inputs.DeviceAuthorizationState;
import java.lang.Boolean;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * The device_authorization resource is used to approve new devices before they can join the tailnet. See https://tailscale.com/kb/1099/device-authorization/ for more details.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.tailscale.TailscaleFunctions;
 * import com.pulumi.tailscale.inputs.GetDeviceArgs;
 * import com.pulumi.tailscale.DeviceAuthorization;
 * import com.pulumi.tailscale.DeviceAuthorizationArgs;
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
 *             .name(&#34;device.example.com&#34;)
 *             .build());
 * 
 *         var sampleAuthorization = new DeviceAuthorization(&#34;sampleAuthorization&#34;, DeviceAuthorizationArgs.builder()        
 *             .deviceId(sampleDevice.applyValue(getDeviceResult -&gt; getDeviceResult.id()))
 *             .authorized(true)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="tailscale:index/deviceAuthorization:DeviceAuthorization")
public class DeviceAuthorization extends com.pulumi.resources.CustomResource {
    /**
     * Whether or not the device is authorized
     * 
     */
    @Export(name="authorized", type=Boolean.class, parameters={})
    private Output<Boolean> authorized;

    /**
     * @return Whether or not the device is authorized
     * 
     */
    public Output<Boolean> authorized() {
        return this.authorized;
    }
    /**
     * The device to set as authorized
     * 
     */
    @Export(name="deviceId", type=String.class, parameters={})
    private Output<String> deviceId;

    /**
     * @return The device to set as authorized
     * 
     */
    public Output<String> deviceId() {
        return this.deviceId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DeviceAuthorization(String name) {
        this(name, DeviceAuthorizationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DeviceAuthorization(String name, DeviceAuthorizationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DeviceAuthorization(String name, DeviceAuthorizationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("tailscale:index/deviceAuthorization:DeviceAuthorization", name, args == null ? DeviceAuthorizationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DeviceAuthorization(String name, Output<String> id, @Nullable DeviceAuthorizationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("tailscale:index/deviceAuthorization:DeviceAuthorization", name, state, makeResourceOptions(options, id));
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
    public static DeviceAuthorization get(String name, Output<String> id, @Nullable DeviceAuthorizationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DeviceAuthorization(name, id, state, options);
    }
}
