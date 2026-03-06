require "./spec_helper"
require "golden"
require "../src/canvas"

{% if flag?(:preview_mt) && flag?(:execution_context) %}
  require "teatest"

  describe "examples/canvas" do
    it "matches golden output for the initial view" do
      output = CanvasModel.new.view.content

      Teatest.wait_for(
        IO::Memory.new(output),
        ->(buf : Bytes) { String.new(buf).includes?(CanvasModel::FOOTER_TEXT) },
        [
          Teatest.with_duration(25.milliseconds),
          Teatest.with_check_interval(1.millisecond),
        ]
      )

      Golden.require_equal("golden/canvas", output)
    end
  end
{% else %}
  describe "examples/canvas" do
    pending "requires -Dpreview_mt -Dexecution_context to enable teatest-based example specs"
  end
{% end %}
