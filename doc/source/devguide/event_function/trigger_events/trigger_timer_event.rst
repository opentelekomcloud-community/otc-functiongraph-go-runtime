Timer Event Source
==================

You can schedule a timer to invoke your code based on a fixed rate of minutes,
hours, or days or a cron expression. For details, see
`Using a Timer Trigger <https://docs.otc.t-systems.com/function-graph/umn/creating_triggers/using_a_timer_trigger.html>`__.

Timer example event
-------------------

.. literalinclude:: /../../samples-doc/event-timer/resources/timer_event.json
    :language: json
    :caption: :github_repo_master:`timer_event.json <samples-doc/event-timer/resources/timer_event.json>`

Parameter description
---------------------

.. list-table::
   :header-rows: 1
   :widths: 20 15 35

   * - Parameter
     - Type
     - Description
   * - version
     - String
     - Event version
   * - time
     - String
     - Time when an event occurs.
   * - trigger_type
     - String
     - Trigger type: **TIMER**
   * - trigger_name
     - String
     - Trigger name
   * - user_event
     - String
     - Additional information of the trigger

Example
-------

.. literalinclude:: /../../samples-doc/event-timer/src/handler.go
    :language: go
    :caption: :github_repo_master:`handler.go <samples-doc/event-timer/src/handler.go>`
