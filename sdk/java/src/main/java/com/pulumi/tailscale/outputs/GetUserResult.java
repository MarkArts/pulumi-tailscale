// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.tailscale.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetUserResult {
    /**
     * @return The time the user joined their tailnet.
     * 
     */
    private String created;
    /**
     * @return true when the user has a node currently connected to the control server.
     * 
     */
    private Boolean currentlyConnected;
    /**
     * @return Number of devices the user owns.
     * 
     */
    private Integer deviceCount;
    /**
     * @return The name of the user.
     * 
     */
    private String displayName;
    /**
     * @return The unique identifier for the user.
     * 
     */
    private @Nullable String id;
    /**
     * @return The later of either: a) The last time any of the user&#39;s nodes were connected to the network or b) The last time the user authenticated to any tailscale service, including the admin panel.
     * 
     */
    private String lastSeen;
    /**
     * @return The emailish login name of the user.
     * 
     */
    private @Nullable String loginName;
    /**
     * @return The profile pic URL for the user.
     * 
     */
    private String profilePicUrl;
    /**
     * @return The role of the user.
     * 
     */
    private String role;
    /**
     * @return The status of the user.
     * 
     */
    private String status;
    /**
     * @return The tailnet that owns the user.
     * 
     */
    private String tailnetId;
    /**
     * @return The type of relation this user has to the tailnet associated with the request.
     * 
     */
    private String type;

    private GetUserResult() {}
    /**
     * @return The time the user joined their tailnet.
     * 
     */
    public String created() {
        return this.created;
    }
    /**
     * @return true when the user has a node currently connected to the control server.
     * 
     */
    public Boolean currentlyConnected() {
        return this.currentlyConnected;
    }
    /**
     * @return Number of devices the user owns.
     * 
     */
    public Integer deviceCount() {
        return this.deviceCount;
    }
    /**
     * @return The name of the user.
     * 
     */
    public String displayName() {
        return this.displayName;
    }
    /**
     * @return The unique identifier for the user.
     * 
     */
    public Optional<String> id() {
        return Optional.ofNullable(this.id);
    }
    /**
     * @return The later of either: a) The last time any of the user&#39;s nodes were connected to the network or b) The last time the user authenticated to any tailscale service, including the admin panel.
     * 
     */
    public String lastSeen() {
        return this.lastSeen;
    }
    /**
     * @return The emailish login name of the user.
     * 
     */
    public Optional<String> loginName() {
        return Optional.ofNullable(this.loginName);
    }
    /**
     * @return The profile pic URL for the user.
     * 
     */
    public String profilePicUrl() {
        return this.profilePicUrl;
    }
    /**
     * @return The role of the user.
     * 
     */
    public String role() {
        return this.role;
    }
    /**
     * @return The status of the user.
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return The tailnet that owns the user.
     * 
     */
    public String tailnetId() {
        return this.tailnetId;
    }
    /**
     * @return The type of relation this user has to the tailnet associated with the request.
     * 
     */
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetUserResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String created;
        private Boolean currentlyConnected;
        private Integer deviceCount;
        private String displayName;
        private @Nullable String id;
        private String lastSeen;
        private @Nullable String loginName;
        private String profilePicUrl;
        private String role;
        private String status;
        private String tailnetId;
        private String type;
        public Builder() {}
        public Builder(GetUserResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.created = defaults.created;
    	      this.currentlyConnected = defaults.currentlyConnected;
    	      this.deviceCount = defaults.deviceCount;
    	      this.displayName = defaults.displayName;
    	      this.id = defaults.id;
    	      this.lastSeen = defaults.lastSeen;
    	      this.loginName = defaults.loginName;
    	      this.profilePicUrl = defaults.profilePicUrl;
    	      this.role = defaults.role;
    	      this.status = defaults.status;
    	      this.tailnetId = defaults.tailnetId;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder created(String created) {
            if (created == null) {
              throw new MissingRequiredPropertyException("GetUserResult", "created");
            }
            this.created = created;
            return this;
        }
        @CustomType.Setter
        public Builder currentlyConnected(Boolean currentlyConnected) {
            if (currentlyConnected == null) {
              throw new MissingRequiredPropertyException("GetUserResult", "currentlyConnected");
            }
            this.currentlyConnected = currentlyConnected;
            return this;
        }
        @CustomType.Setter
        public Builder deviceCount(Integer deviceCount) {
            if (deviceCount == null) {
              throw new MissingRequiredPropertyException("GetUserResult", "deviceCount");
            }
            this.deviceCount = deviceCount;
            return this;
        }
        @CustomType.Setter
        public Builder displayName(String displayName) {
            if (displayName == null) {
              throw new MissingRequiredPropertyException("GetUserResult", "displayName");
            }
            this.displayName = displayName;
            return this;
        }
        @CustomType.Setter
        public Builder id(@Nullable String id) {

            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder lastSeen(String lastSeen) {
            if (lastSeen == null) {
              throw new MissingRequiredPropertyException("GetUserResult", "lastSeen");
            }
            this.lastSeen = lastSeen;
            return this;
        }
        @CustomType.Setter
        public Builder loginName(@Nullable String loginName) {

            this.loginName = loginName;
            return this;
        }
        @CustomType.Setter
        public Builder profilePicUrl(String profilePicUrl) {
            if (profilePicUrl == null) {
              throw new MissingRequiredPropertyException("GetUserResult", "profilePicUrl");
            }
            this.profilePicUrl = profilePicUrl;
            return this;
        }
        @CustomType.Setter
        public Builder role(String role) {
            if (role == null) {
              throw new MissingRequiredPropertyException("GetUserResult", "role");
            }
            this.role = role;
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            if (status == null) {
              throw new MissingRequiredPropertyException("GetUserResult", "status");
            }
            this.status = status;
            return this;
        }
        @CustomType.Setter
        public Builder tailnetId(String tailnetId) {
            if (tailnetId == null) {
              throw new MissingRequiredPropertyException("GetUserResult", "tailnetId");
            }
            this.tailnetId = tailnetId;
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            if (type == null) {
              throw new MissingRequiredPropertyException("GetUserResult", "type");
            }
            this.type = type;
            return this;
        }
        public GetUserResult build() {
            final var _resultValue = new GetUserResult();
            _resultValue.created = created;
            _resultValue.currentlyConnected = currentlyConnected;
            _resultValue.deviceCount = deviceCount;
            _resultValue.displayName = displayName;
            _resultValue.id = id;
            _resultValue.lastSeen = lastSeen;
            _resultValue.loginName = loginName;
            _resultValue.profilePicUrl = profilePicUrl;
            _resultValue.role = role;
            _resultValue.status = status;
            _resultValue.tailnetId = tailnetId;
            _resultValue.type = type;
            return _resultValue;
        }
    }
}
