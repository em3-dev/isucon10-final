# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: xsuportal/services/registration/session.proto

require 'google/protobuf'

require 'xsuportal/resources/team_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("xsuportal/services/registration/session.proto", :syntax => :proto3) do
    add_message "xsuportal.proto.services.registration.GetRegistrationSessionQuery" do
      optional :team_id, :int64, 1
      optional :invite_token, :string, 2
    end
    add_message "xsuportal.proto.services.registration.GetRegistrationSessionResponse" do
      optional :team, :message, 1, "xsuportal.proto.resources.Team"
      optional :status, :enum, 2, "xsuportal.proto.services.registration.GetRegistrationSessionResponse.Status"
      optional :member_invite_url, :string, 3
      optional :invite_token, :string, 4
    end
    add_enum "xsuportal.proto.services.registration.GetRegistrationSessionResponse.Status" do
      value :CLOSED, 0
      value :NOT_JOINABLE, 1
      value :NOT_LOGGED_IN, 2
      value :CREATABLE, 3
      value :JOINABLE, 4
      value :JOINED, 5
    end
    add_message "xsuportal.proto.services.registration.UpdateRegistrationRequest" do
      optional :team_name, :string, 1
      optional :name, :string, 2
      optional :email_address, :string, 3
      optional :is_student, :bool, 4
    end
    add_message "xsuportal.proto.services.registration.UpdateRegistrationResponse" do
    end
    add_message "xsuportal.proto.services.registration.DeleteRegistrationRequest" do
    end
    add_message "xsuportal.proto.services.registration.DeleteRegistrationResponse" do
    end
  end
end

module Xsuportal
  module Proto
    module Services
      module Registration
        GetRegistrationSessionQuery = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("xsuportal.proto.services.registration.GetRegistrationSessionQuery").msgclass
        GetRegistrationSessionResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("xsuportal.proto.services.registration.GetRegistrationSessionResponse").msgclass
        GetRegistrationSessionResponse::Status = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("xsuportal.proto.services.registration.GetRegistrationSessionResponse.Status").enummodule
        UpdateRegistrationRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("xsuportal.proto.services.registration.UpdateRegistrationRequest").msgclass
        UpdateRegistrationResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("xsuportal.proto.services.registration.UpdateRegistrationResponse").msgclass
        DeleteRegistrationRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("xsuportal.proto.services.registration.DeleteRegistrationRequest").msgclass
        DeleteRegistrationResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("xsuportal.proto.services.registration.DeleteRegistrationResponse").msgclass
      end
    end
  end
end
