#!/usr/bin/env sh

su - sa_cluster

spadmin external_view external_dimension_table add \
-p default \
-t items \
-e "events.issue_id=items.item_id AND items.item_type = 'issue'"

spadmin external_view external_dimension_table add \
-p default \
-t items#1 \
-e "events.meeting_id=items#1.item_id AND items#1.item_type = 'meeting'"

spadmin external_view external_dimension_table add \
-p default \
-t items#2 \
-e "events.business_id=items#2.item_id AND events.business_type='事项' AND items#2.item_type = 'issue'"

spadmin external_view external_dimension_table add \
-p default \
-t items#3 \
-e "events.business_id=items#3.item_id AND events.business_type='会议' AND items#3.item_type = 'meeting'"

spadmin external_view external_dimension_table add \
-p default \
-t items#4 \
-e "events.project_id=items#4.item_id AND items#4.item_type = 'project'"

spadmin external_view external_dimension_table add \
-p default \
-t items#5 \
-e "events.business_id=items#5.item_id AND events.business_type='项目' AND items#5.item_type = 'project'"

spadmin external_view external_dimension_table add \
-p default \
-t items#6 \
-e "events.project_id=items#6.item_id AND items#6.item_type = 'workspace'"

spadmin external_view external_dimension_table add \
-p default \
-t items#7 \
-e "events.business_id=items#7.item_id AND events.business_type='空间' AND items#7.item_type = 'workspace'"
